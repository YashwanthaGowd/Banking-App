package main

import (
	"log"

	"github.com/banking-app/account-service/src/config"
	"github.com/banking-app/account-service/src/gateway"
	"github.com/banking-app/account-service/src/handler"
	"github.com/banking-app/account-service/src/server"
	bankingService "github.com/banking-app/account-service/src/service/banking"
	kafkaService "github.com/banking-app/account-service/src/service/kafka"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := fx.New(
		// Provide all the constructors
		fx.Provide(
			config.LoadFromFile,
			bankingService.NewService,
			kafkaService.NewKafkaService,
			gateway.NewGateway,
			handler.NewHandler,
			server.NewGinServer,
		),
		// Invoke runs the application
		fx.Invoke(
			server.RunServer,
			kafkaService.StartKafkaScan,
		),
	)

	app.Run()
}


