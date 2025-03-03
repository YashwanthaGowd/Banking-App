package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/banking-app/transaction-service/src/config"
	"github.com/banking-app/transaction-service/src/model"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TransactionService interface {
	GetTransactionsbyMonthRange(accountId string, startMonth time.Time, endMonth time.Time) ([]model.Transaction, error)
	GetTransactionsbyCount(accountId string, count int) ([]model.Transaction, error)
	GetTransactionbyId(TransactionId string) (model.Transaction, error)
	AddTransaction(transaction *model.Transaction) (string, error)
}

type transactionService struct {
	db *mongo.Database
}

func NewTransactionService(cfg *config.Config) (TransactionService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(cfg.MongoDB.URI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database("banking")

	log.Println("Connected to MongoDB")

	return &transactionService{db: db}, nil
}

func (ts *transactionService) AddTransaction(transaction *model.Transaction) (string, error) {
	collection := ts.db.Collection("transactions")
	transaction.ID = uuid.New().String()
	transaction.Timestamp = time.Now()
	_, err := collection.InsertOne(context.Background(), transaction)
	if err != nil {
		return "", err
	}
	return transaction.ID, nil
}

func (ts *transactionService) GetTransactionbyId(id string) (model.Transaction, error) {
	collection := ts.db.Collection("transactions")
	var transaction model.Transaction
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.Background(), filter).Decode(&transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	if transaction.ID == "" {
		return model.Transaction{}, fmt.Errorf("transaction not found")
	}	

	return transaction, nil
}


func (ts *transactionService) GetTransactionsbyMonthRange(accountId string, startMonth time.Time, endMonth time.Time) ([]model.Transaction, error) {
	
	if endMonth.Before(startMonth) {
		return nil, fmt.Errorf("end month cannot be less than start month")
	}

	collection := ts.db.Collection("transactions")
	filter := bson.M{"account": accountId, "timestamp": bson.M{"$gte": startMonth, "$lt": endMonth}}
	var transactions []model.Transaction
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var transaction model.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transactions found")
	}

	return transactions, nil
}

func (ts *transactionService) GetTransactionsbyCount(accountId string, count int) ([]model.Transaction, error) {

	collection := ts.db.Collection("transactions")
	filter := bson.M{"account": accountId}

	opts := options.Find().SetSort(bson.D{{Key: "timestamp", Value: -1}}).SetLimit(30)

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())
	var transactions []model.Transaction
	for cursor.Next(context.Background()) {
		var transaction model.Transaction
		err := cursor.Decode(&transaction)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}


	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transactions found")
	}

	return transactions, nil
}
