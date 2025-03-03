package server

import (
	"context"
	"log"

	"github.com/banking-app/account-service/src/config"
	"github.com/banking-app/account-service/src/handler"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func NewGinServer(accountHandler handler.Handler) *gin.Engine {
	r := gin.Default()
	bankingApp := r.Group("/bankingapp")
	accountGroup := bankingApp.Group("/accounts")

	accountGroup.POST("", accountHandler.CreateAccount)
	accountGroup.GET("/:accountId", accountHandler.GetAccountbyId)
	accountGroup.PUT("/:accountId", accountHandler.UpdateAccount)
	accountGroup.DELETE("/:accountId", accountHandler.DisableAccount)
	accountGroup.PATCH("/:accountId", accountHandler.ActivateAccount)
	accountGroup.POST("/deposit", accountHandler.Deposit)
	accountGroup.POST("/withdraw", accountHandler.Withdraw)
	accountGroup.GET("/transactions/history/:account/:count", accountHandler.GetTransactionsbyAccount)
	accountGroup.GET("/transactions/range/:account/:startMonth/:endMonth", accountHandler.GetTransactionsbyMonthRange)
	accountGroup.GET("/transactions/id/:transactionId", accountHandler.GetTransactionbyId)
	

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