package service

import (

	//import config from src/config/config.go

	"context"
	"database/sql"
	"fmt"

	"github.com/banking-app/account-service/src/model"

	"github.com/lib/pq"
)

func (s *bankingService) CreateTransaction(transaction *model.Transaction) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()
	res, err := tx.Exec("INSERT INTO transactions (id, account, amount, type, timestamp) VALUES ($1, $2, $3, $4, $5)", transaction.ID, transaction.Account, transaction.Amount, transaction.Type, transaction.Timestamp)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("transaction not created")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (s *bankingService) GetTransactions() ([]model.Transaction, error) {
	rows, err := s.db.Query("SELECT id, account, amount, type, timestamp FROM transactions")
	if err != nil {
		return nil, fmt.Errorf("failed to query transactions: %v", err)
	}
	defer rows.Close()

	var transactions []model.Transaction
	for rows.Next() {
		var t model.Transaction
		err := rows.Scan(&t.ID, &t.Account, &t.Amount, &t.Type, &t.Timestamp)
		if err != nil {
			return nil, fmt.Errorf("failed to scan transaction: %v", err)
		}
		transactions = append(transactions, t)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating transactions: %v", err)
	}
	return transactions, nil
}

func (s *bankingService) DeleteTransactionsByIds(ids []string) error {
	tx, err := s.db.Begin()

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Execute delete within transaction
	res, err := tx.Exec(`
		DELETE FROM transactions 
		WHERE id IN ($1)`,
		pq.Array(ids))

	if err != nil {
		return fmt.Errorf("failed to delete transactions: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("transactions not deleted")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}

func (s *bankingService) DeleteTransactionsById(id string) error {
	tx, err := s.db.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()
	res, err := tx.Exec(`
		DELETE FROM transactions 
		WHERE id = $1`,
		id)

	if err != nil {
		return fmt.Errorf("failed to delete transactions: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("transaction not deleted")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}
	return nil
}
