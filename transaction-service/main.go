package main

import (
	"context"
	"log"

	"github.com/banking-app/transaction-service/src/config"
	"github.com/banking-app/transaction-service/src/handler"
	"github.com/banking-app/transaction-service/src/server"
	kafkaservice "github.com/banking-app/transaction-service/src/service/kafka"
	transactionService "github.com/banking-app/transaction-service/src/service/transaction"
	"go.uber.org/fx"
)

func main() {
	// Initialize your config and transaction service

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := fx.New(
		// Provide all the constructors
		fx.Provide(
			config.LoadFromFile,
			transactionService.NewTransactionService,
			kafkaservice.NewKafkaConsumer,
			handler.NewHandler,
			server.NewGinServer,
		),
		// Invoke runs the application
		fx.Invoke(
			server.RunServer,
			startKafkaConsumer,
		),
	)

	app.Run()
}

func startKafkaConsumer(lc fx.Lifecycle, consumer *kafkaservice.KafkaConsumer) {
	ctx := context.Background()
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go kafkaservice.StartConsuming(ctx, consumer)
			return nil
		},
		OnStop: func(context.Context) error {
			// Add cleanup if necessary
			return nil
		},
	})
}
