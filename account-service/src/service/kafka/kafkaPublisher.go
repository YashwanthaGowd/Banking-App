package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/banking-app/account-service/src/config"
	"github.com/banking-app/account-service/src/model"
	bankingService "github.com/banking-app/account-service/src/service/banking"
	"go.uber.org/fx"

	"github.com/segmentio/kafka-go"
)

type KafkaService interface {
	PublishTransaction(transaction *model.Transaction) error
}

type kafkaService struct {
	banking bankingService.BankingService
	writer *kafka.Writer
	topic  string
}

func NewKafkaService(cfg *config.Config, banking bankingService.BankingService) (KafkaService, error) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      cfg.Kafka.Brokers,
		Topic:        cfg.Kafka.Topic,
		BatchSize:    cfg.Kafka.BatchSize,
		BatchTimeout: time.Millisecond * time.Duration(cfg.Kafka.BatchTimeout),
		RequiredAcks: cfg.Kafka.RequiredAcks,
		Async:        cfg.Kafka.Async,
	})

	return &kafkaService{
		banking: banking,
		writer: writer,
		topic:  "banking.transactions",
	}, nil
}

func (k *kafkaService) PublishTransaction(transaction *model.Transaction) error {
	json, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Value: json,
		Headers: []kafka.Header{
			{
				Key:   "type",
				Value: []byte(transaction.ID),
			},
		},
	}

	ctx := context.Background()
	err = k.writer.WriteMessages(ctx, msg)
	if err != nil {
		log.Printf("Failed to publish transaction: %v", err)
		return err
	}

	return nil
}

// Don't forget to add a Close method if you need to clean up
func Close(k *kafkaService) error {
	return k.writer.Close()
}

// scan the transactions and publish to kafka periodically using a ticker in for loop
func(kafkaservice *kafkaService) ScanTransactions() error {
	t := time.NewTicker(time.Second * 10)
	// Scan transactions every 10 seconds
	for {
		select {
		case <-t.C:
			
			// get transactions from postgres
			transactions, err := kafkaservice.banking.GetTransactions()
			if err != nil {
				log.Printf("Error getting transactions: %v", err)
				continue
			}
			for _, transaction := range transactions {
				err:= kafkaservice.PublishTransaction(&transaction)
				if err != nil {
					log.Printf("Error publishing transaction: %v", err)
					continue
				}
				log.Printf("Successfully published transaction: %v", transaction)
				err= kafkaservice.banking.DeleteTransactionsById(transaction.ID)
				if err != nil {
					log.Printf("Error deleting transaction: %v", err)
					continue
				}
			}

		}
	}
}

func StartKafkaScan(lc fx.Lifecycle, kafka KafkaService) {
    lc.Append(fx.Hook{
        OnStart: func(context.Context) error {
            go func() {
                if k, ok := kafka.(*kafkaService); ok {
                    if err := k.ScanTransactions(); err != nil {
                        log.Printf("Error starting transaction scan: %v", err)
                    }
                }
            }()
            return nil
        },
        OnStop: func(context.Context) error {
            Close(kafka.(*kafkaService))
            return nil
        },
    })
}

