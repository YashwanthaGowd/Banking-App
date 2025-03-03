package handler

import (
	"fmt"
	"net/http"

	accountpb "github.com/banking-app/protos/generated/account"
	"github.com/google/uuid"

	"github.com/banking-app/account-service/src/model"

	"github.com/gin-gonic/gin"
)

// createAccount creates a new account
func (h handler) CreateAccount(c *gin.Context) {
	req := &accountpb.CreateAccountRequest{}

	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := model.NewAccountFromProto(req)
	err := h.BankingService.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// publish opening balance
	err = h.KafkaService.PublishTransaction(model.NewTransaction(account.ID, account.Balance, "opening"))
	if err != nil {
		err= h.BankingService.CreateTransaction(model.NewTransaction(account.ID, account.Balance, "opening"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Account %s created successfully", account.ID)})
}

// GetAccountbyId gets a account by id
func (h handler) GetAccountbyId(c *gin.Context) {
	accountId := c.Param("accountId")
	account, err := h.BankingService.GetAccountbyId(accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, account)
}

func (h handler) UpdateAccount(c *gin.Context) {
	req := &accountpb.UpdateAccountRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if account exists
	account, err := h.BankingService.GetAccountbyId(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	// update account
	account.FirstName = req.FirstName
	account.LastName = req.LastName
	account.Email = req.Email
	account.Type = req.Type
	account.Password = req.Password

	err = h.BankingService.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account updated successfully"})
}

// disableAccount disables an account
func (h handler) DisableAccount(c *gin.Context) {
	accountId := c.Param("accountId")
	//  get account
	account, err := h.BankingService.GetAccountbyId(accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	// update account
	if account.Status == "closed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account is already closed"})
		return
	}
	account.Status = "closed"

	err = h.BankingService.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account Disabled successfully"})
}

// activateAccount activates an account
func (h handler) ActivateAccount(c *gin.Context) {
	accountId := c.Param("accountId")

	//  get account
	account, err := h.BankingService.GetAccountbyId(accountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	// update account
	if account.Status == "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account is already active"})
		return
	}
	account.Status = "active"

	err = h.BankingService.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Account Activated successfully"})
}

func (h handler) Deposit(c *gin.Context) {

	req := &accountpb.DepositRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if account exists
	account, err := h.BankingService.GetAccountbyId(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	// update account
	account.Balance = account.Balance + req.Amount

	err = h.BankingService.Deposit(account.ID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// publish deposit
	err = h.KafkaService.PublishTransaction(model.NewTransaction(account.ID, req.Amount, "credit"))
	if err != nil {
		err= h.BankingService.CreateTransaction(model.NewTransaction(account.ID, req.Amount, "credit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s has deposited %f", account.ID, req.Amount)})

}

func (h handler) Withdraw(c *gin.Context) {

	req := &accountpb.WithdrawRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// find if account exists
	account, err := h.BankingService.GetAccountbyId(req.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if account == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "account not found"})
		return
	}
	// update account
	account.Balance = account.Balance - req.Amount

	err = h.BankingService.Withdraw(account.ID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// publish withdrawal
	err = h.KafkaService.PublishTransaction(model.NewTransaction(account.ID, req.Amount, "debit"))
	if err != nil {
		err= h.BankingService.CreateTransaction(model.NewTransaction(account.ID, req.Amount, "debit"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s has withdrawn %f", account.ID, req.Amount)})

}

func (h handler) PublishTransactions(c *gin.Context) {
	err:= h.KafkaService.PublishTransaction(model.NewTransaction(uuid.New().String(), 0, "opening"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction published successfully"})
}