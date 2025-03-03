package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/banking-app/transaction-service/src/config"
	"github.com/banking-app/transaction-service/src/model"
	service "github.com/banking-app/transaction-service/src/service/transaction"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	config     config.Kafka
	txnService service.TransactionService
}

func NewKafkaConsumer(config *config.Config, txnService service.TransactionService) *KafkaConsumer {
	return &KafkaConsumer{
		config:     config.Kafka,
		txnService: txnService,
	}
}

func StartConsuming(ctx context.Context, kafkaConsumer *KafkaConsumer) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kafkaConsumer.config.Brokers,
		Topic:   kafkaConsumer.config.Topic,
		GroupID: kafkaConsumer.config.ConsumerGroup,
	})
	defer reader.Close()

	log.Printf("Started consuming from topic: banking.transactions")

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Printf("Error reading message: %v", err)
				continue
			}

			var transaction model.Transaction
			if err := json.Unmarshal(msg.Value, &transaction); err != nil {
				log.Printf("Failed to unmarshal transaction: %v", err)
				continue
			}

			// Add transaction to MongoDB
			id, err := kafkaConsumer.txnService.AddTransaction(&transaction)
			if err != nil {
				log.Printf("Failed to store transaction: %v", err)
				continue
			}

			log.Printf("Successfully processed transaction with ID: %s", id)
		}
	}
}
