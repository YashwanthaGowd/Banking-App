package model

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        string    `json:"id" bson:"_id"`
	Account   string    `json:"account" bson:"account"`
	Amount    float64   `json:"amount" bson:"amount"`
	Type      string    `json:"type" bson:"type"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
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
