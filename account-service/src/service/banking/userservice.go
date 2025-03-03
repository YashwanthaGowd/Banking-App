package service

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/banking-app/account-service/src/model"
)

// User methods
func (s *bankingService) GetUserbyEmail(userId string) (*model.User, error) {
	var user model.User
	err := s.db.QueryRow("SELECT * FROM users WHERE email = $1", userId).Scan(&user.FirstName, &user.LastName, &user.Email, &user.Type, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (s *bankingService) CreateUser(user *model.User) error {
	
	tx, err := s.db.Begin()

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer tx.Rollback()

	// Execute insert within transaction
	res, err := tx.Exec(`
		INSERT INTO users (
			first_name, last_name, email, type, 
			password, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		user.FirstName, user.LastName, user.Email, user.Type,
		user.Password, user.CreatedAt, user.UpdatedAt)
		
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not created")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *bankingService) UpdateUser(user *model.User) error {

	tx, err := s.db.Begin()

	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer tx.Rollback()

	// Execute update within transaction
	res, err := tx.Exec(`
		UPDATE users 
		SET first_name = $1, last_name = $2, type = $3, 
			password = $4, updated_at = $5 
		WHERE email = $6`,
		user.FirstName, user.LastName, user.Type,	
		user.Password, user.UpdatedAt, user.Email)
		
	if err != nil {	
		return fmt.Errorf("failed to update user: %v", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not updated")
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil	
}
