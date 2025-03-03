package model

import (
	"time"

	accountpb "github.com/banking-app/protos/generated/account"

	"github.com/google/uuid"
)

type Account struct {
	ID        string    `json:"id" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Type      string    `json:"type" db:"account_type"`
	Balance   float64   `json:"balance" db:"balance"`
	Status    string    `json:"status" db:"status"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type User struct {
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Type      string    `json:"type" db:"type"`
	Status    string    `json:"status" db:"status"`
	Password  string    `json:"password" db:"password"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

func (a *Account) ToProto() *accountpb.Account {
	return &accountpb.Account{
		Id:          a.ID,
		FirstName:   a.FirstName,
		LastName:    a.LastName,
		Email:       a.Email,
		AccountType: a.Type,
		Balance:     a.Balance,
		Status:      a.Status,
		Password:    a.Password,
	}
}

func NewAccountFromProto(a *accountpb.CreateAccountRequest) *Account {
	return &Account{
		ID:        uuid.New().String(),
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Email:     a.Email,
		Type:      a.Type,
		Balance:   a.Balance,
		Status:    a.Status,
		Password:  a.Password,
	}
}

func NewUserFromProto(u *accountpb.CreateUserRequest) *User {
	return &User{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Type:      u.Type,
		Password:  u.Password,
	}
}
