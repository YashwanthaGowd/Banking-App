package handler

import (
	"github.com/banking-app/account-service/src/gateway"
	bankingService "github.com/banking-app/account-service/src/service/banking"
	kafkaService "github.com/banking-app/account-service/src/service/kafka"

	"github.com/gin-gonic/gin"
)

type Handler interface {

	// Account methods
	CreateAccount(c *gin.Context)
	GetAccountbyId(c *gin.Context)
	UpdateAccount(c *gin.Context)
	DisableAccount(c *gin.Context)
	ActivateAccount(c *gin.Context)
	Deposit(c *gin.Context)
	Withdraw(c *gin.Context)

	// User methods 
	CreateUser(c *gin.Context)
	GetUserbyEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	DisableUser(c *gin.Context)
	ActivateUser(c *gin.Context)

	// Transaction methods
	GetTransactionbyId(c *gin.Context)
	GetTransactionsbyAccount(c *gin.Context)
	GetTransactionsbyMonthRange(c *gin.Context)
}

// AccountHandlerImpl implements AccountHandler
type handler struct {
	BankingService bankingService.BankingService
	KafkaService   kafkaService.KafkaService
	Gateway        gateway.Gateway
}

// NewAccountHandlerImpl returns a new AccountHandlerImpl
func NewHandler(bankingService bankingService.BankingService, kafkaService kafkaService.KafkaService, gateway gateway.Gateway) Handler {
	return &handler{
		BankingService: bankingService,
		KafkaService:   kafkaService,
		Gateway:        gateway,
	}
}
