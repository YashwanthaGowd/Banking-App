package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/banking-app/account-service/src/model"
)

// Account methods
func (s *bankingService) GetAccountbyId(accountId string) (*model.Account, error) {
	var account model.Account
	err := s.db.QueryRow("SELECT * FROM accounts WHERE id = $1", accountId).Scan(
		&account.ID, &account.FirstName, &account.LastName, &account.Email,
		&account.Type, &account.Balance, &account.Status, &account.Password,
		&account.CreatedAt, &account.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("account not found")
		}
		return nil, err
	}
	return &account, nil
}

func (s *bankingService) CreateAccount(account *model.Account) error {
	// Start transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer tx.Rollback()

	// Execute insert within transaction
	res, err := tx.Exec(`
		INSERT INTO accounts (
			id, first_name, last_name, email, account_type, 
			balance, status, password, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`,
		account.ID, account.FirstName, account.LastName, account.Email,
		account.Type, account.Balance, account.Status, account.Password,
		account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert account: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("account not created")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}


	return nil
}

func (s *bankingService) UpdateAccount(account *model.Account) error {
	// Start transaction
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// First verify account exists
	var currentBalance float64
	err = tx.QueryRow("SELECT balance FROM accounts WHERE id = $1 FOR UPDATE", account.ID).Scan(&currentBalance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("account not found")
		}
		return fmt.Errorf("failed to query account: %v", err)
	}

	// Execute update within transaction
	res, err := tx.Exec(`
		UPDATE accounts 
		SET first_name = $1, last_name = $2, email = $3, 
			account_type = $4, balance = $5, status = $6, 
			password = $7, updated_at = $8 
		WHERE id = $9`,
		account.FirstName, account.LastName, account.Email,
		account.Type, account.Balance, account.Status,
		account.Password, account.UpdatedAt, account.ID)
	if err != nil {
		return fmt.Errorf("failed to update account: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("account not found")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

// Deposit amount to an account
func (s *bankingService) Deposit(accountID string, amount float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get current balance with row lock
	var account model.Account
	err = tx.QueryRow(`
		SELECT id, balance, status 
		FROM accounts 
		WHERE id = $1 
		FOR UPDATE`, accountID).Scan(&account.ID, &account.Balance, &account.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("account not found")
		}
		return fmt.Errorf("failed to query account: %v", err)
	}

	if account.Status != "active" {
		return fmt.Errorf("account is not active")
	}

	// Update balance
	newBalance := account.Balance + amount
	_, err = tx.Exec(`
		UPDATE accounts 
		SET balance = $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2`,
		newBalance, accountID)
	if err != nil {
		return fmt.Errorf("failed to update balance: %v", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *bankingService) Withdraw(accountID string, amount float64) error {
	tx, err := s.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	// Get current balance with row lock
	var account model.Account
	err = tx.QueryRow(`
		SELECT id, balance, status 
		FROM accounts 
		WHERE id = $1 
		FOR UPDATE`, accountID).Scan(&account.ID, &account.Balance, &account.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("account not found")
		}
		return fmt.Errorf("failed to query account: %v", err)
	}

	if account.Status != "active" {
		return fmt.Errorf("account is not active")
	}

	if account.Balance < amount {
		return fmt.Errorf("insufficient funds")
	}

	// Update balance
	newBalance := account.Balance - amount
	_, err = tx.Exec(`
		UPDATE accounts 
		SET balance = $1, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $2`,
		newBalance, accountID)
	if err != nil {
		return fmt.Errorf("failed to update balance: %v", err)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}