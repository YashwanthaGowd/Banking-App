package service

import (
	"testing"
	"time"

	"github.com/banking-app/account-service/src/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockTransactionService struct {
	t *testing.T
	mock.Mock
}

func (m *mockTransactionService) CreateTransaction(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *mockTransactionService) GetTransactions() ([]model.Transaction, error) {
	args := m.Called()
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func (m *mockTransactionService) DeleteTransactionsByIds(ids []string) error {
	args := m.Called(ids)
	return args.Error(0)
}

func (m *mockTransactionService) DeleteTransactionsById(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func NewMockTransactionService(t *testing.T) *mockTransactionService {
	return &mockTransactionService{t: t}
}

func TestCreateTransaction(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("CreateTransaction", &transaction).Return(nil)

	// call the CreateTransaction method
	err := mockTransactionService.CreateTransaction(&transaction)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestGetTransactions(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("GetTransactions").Return([]model.Transaction{transaction}, nil)

	// call the GetTransactions method
	result, err := mockTransactionService.GetTransactions()

	if len(result) != 1 {
		t.Errorf("Expected result to have 1 element, but got %d", len(result))
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestDeleteTransactionsByIds(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("DeleteTransactionsByIds", []string{transaction.ID}).Return(nil)

	// call the DeleteTransactionsByIds method
	err := mockTransactionService.DeleteTransactionsByIds([]string{transaction.ID})

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestDeleteTransactionsById(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("DeleteTransactionsById", transaction.ID).Return(nil)

	// call the DeleteTransactionsById method
	err := mockTransactionService.DeleteTransactionsById(transaction.ID)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

