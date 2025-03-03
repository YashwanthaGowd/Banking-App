package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        string    `json:"id"`
	Account    string    `json:"account"`
	Amount    float64       `json:"amount"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

func NewTransaction(Account string, Amount float64, Type string) *Transaction {
	return &Transaction{
		ID:        uuid.New().String(),
		Account:   Account,
		Amount:    Amount,	
		Type:      Type,
		Timestamp: time.Now(),
	}
}
