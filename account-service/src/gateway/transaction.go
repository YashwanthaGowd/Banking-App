package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/banking-app/account-service/src/model"
)

type gateway struct {
	transactionClient *http.Client
}

type Gateway interface {
	GetTransactionbyId(transactionId string) (model.Transaction, error)
	GetTransactionsbyAccount(accountId string, count int) ([]model.Transaction, error)
	GetTransactionsbyMonthRange(accountId string, startMonth string, endMonth string) ([]model.Transaction, error)
}

func NewGateway() Gateway {
	client := &http.Client{}
	return &gateway{
		transactionClient: client,
	}
}


func (g *gateway) GetTransactionbyId(transactionId string) (model.Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/bankingapp/transactions/id/%s", transactionId), nil)
	if err != nil {
		return model.Transaction{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := g.transactionClient.Do(req)
	if err != nil {
		return model.Transaction{}, err
	}

	if res.StatusCode != http.StatusOK {
		return model.Transaction{}, fmt.Errorf("no transactions found")
	}
	defer res.Body.Close()
	
	transaction, err := g.parseTransaction(res.Body)
	if err != nil {
		return model.Transaction{}, err
	}
	if transaction.ID == "" {
		return model.Transaction{}, fmt.Errorf("no transactions found")
	}
	return transaction, nil
}


func (g *gateway) GetTransactionsbyAccount(accountId string, count int) ([]model.Transaction, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/bankingapp/transactions/history/%s/%d", accountId, count), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := g.transactionClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("no transactions found")
	}

	defer res.Body.Close()
	transactions, err := g.parseTransactions(res.Body)
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transactions found")
	}

	return transactions, nil
}


func (g *gateway) GetTransactionsbyMonthRange(accountId string, startMonth string, endMonth string) ([]model.Transaction, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("http://localhost:8081/bankingapp/transactions/range/%s/%s/%s", accountId, startMonth, endMonth), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := g.transactionClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("no transactions found")
	}
	defer res.Body.Close()

	transactions, err := g.parseTransactions(res.Body)
	if err != nil {
		return nil, err
	}

	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transactions found")
	}
	return transactions, nil
}

func (g *gateway) parseTransactions(body io.ReadCloser) ([]model.Transaction, error) {
	var transactions []model.Transaction
	decoder := json.NewDecoder(body)
	for decoder.More() {
		err := decoder.Decode(&transactions)
		if err != nil {
			return nil, err
		}		
	}
	return transactions, nil
}
func (g *gateway) parseTransaction(body io.ReadCloser) (model.Transaction, error) {
	var transactions model.Transaction
	decoder := json.NewDecoder(body)
	for decoder.More() {
		var transaction model.Transaction
		err := decoder.Decode(&transaction)
		if err != nil {
			return model.Transaction{}, err
		}

		if transaction.ID == "" {
			return model.Transaction{}, fmt.Errorf("transaction not found")
		}
		
	}
	return transactions, nil
}
