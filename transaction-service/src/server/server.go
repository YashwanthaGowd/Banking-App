package server

import (
	"context"
	"log"

	"github.com/banking-app/transaction-service/src/config"
	"github.com/banking-app/transaction-service/src/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinServer(handler handler.Handler) *gin.Engine {
	r := gin.Default()
	bankingApp := r.Group("/bankingapp")

	transactionGroup := bankingApp.Group("/transactions")
	transactionGroup.GET("/id/:transactionId", handler.GetTransactionbyId)
	transactionGroup.GET("history/:account/:count", handler.GetTransactionsbyCount)
	transactionGroup.GET("/range/:account/:startMonth/:endMonth", handler.GetTransactionsbyMonthRange)

	return r

}

// runServer starts the Gin server
func RunServer(lc fx.Lifecycle, ginServer *gin.Engine, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Printf("Starting server on %s:%s", cfg.Server.Host, cfg.Server.Port)
			go ginServer.Run(cfg.Server.Host + ":" + cfg.Server.Port)
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Shutting down server")
			return nil
		},
	})
}
