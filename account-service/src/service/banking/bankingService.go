package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/banking-app/account-service/src/config"
	"github.com/banking-app/account-service/src/model"

	_ "github.com/lib/pq"
)

type BankingService interface {
	// Account methods
	GetAccountbyId(accountId string) (*model.Account, error)
	CreateAccount(account *model.Account) error
	UpdateAccount(account *model.Account) error
	Deposit(accountID string, amount float64) error
	Withdraw(accountID string, amount float64) error

	// User methods
	GetUserbyEmail(userId string) (*model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user *model.User) error

	// Transaction methods
	CreateTransaction(transaction *model.Transaction) error
	GetTransactions() ([]model.Transaction, error)
	DeleteTransactionsById(id string) error
	DeleteTransactionsByIds(ids []string) error
	// PublishTransaction(transaction *model.Transaction) error
}

type bankingService struct {
	db *sql.DB
}

func NewService(cfg *config.Config) (BankingService, error) {
	db, err := sql.Open("postgres", cfg.Postgres.Uri)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	schemaSQL, err := os.ReadFile("resources/000001_create_tables.sql")
	if err != nil {
		return nil, fmt.Errorf("error reading schema file: %v", err)
	}

	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		return nil, fmt.Errorf("error creating schema: %v", err)
	}

	return &bankingService{
		db: db,
	}, nil
}
