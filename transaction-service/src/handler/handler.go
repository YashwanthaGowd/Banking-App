package handler

import (
	transactionservice "github.com/banking-app/transaction-service/src/service/transaction"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetTransactionbyId(c *gin.Context)
	GetTransactionsbyCount(c *gin.Context)
	GetTransactionsbyMonthRange(c *gin.Context)
}

// AccountHandlerImpl implements AccountHandler
type handler struct {
	TransactionService transactionservice.TransactionService
}

// NewAccountHandlerImpl returns a new AccountHandlerImpl
func NewHandler(transactionService transactionservice.TransactionService) Handler {
	return &handler{
		TransactionService: transactionService,
	}
}
