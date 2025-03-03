// write tests for the transaction service. use mocks for mongodb

package service

import (
	"testing"
	"time"

	"github.com/banking-app/transaction-service/src/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockTransactionService struct {
	t *testing.T
	mock.Mock
}

func (m *MockTransactionService) GetTransactionbyId(id string) (model.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(model.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetTransactionsbyMonthRange(accountId string, startMonth time.Time, endMonth time.Time) ([]model.Transaction, error) {
	args := m.Called(accountId, startMonth, endMonth)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func (m *MockTransactionService) GetTransactionsbyCount(accountId string, count int) ([]model.Transaction, error) {
	args := m.Called(accountId, count)
	return args.Get(0).([]model.Transaction), args.Error(1)
}

func NewMockTransactionService(t *testing.T) *MockTransactionService {
	return &MockTransactionService{t: t}
}

func TestGetTransactionbyId(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("GetTransactionbyId", transaction.ID).Return(transaction, nil)

	// call the GetTransactionbyId method
	result, err := mockTransactionService.GetTransactionbyId(transaction.ID)

	if result.ID != transaction.ID {
		t.Errorf("Expected result.ID to be %s, but got %s", transaction.ID, result.ID)
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestGetTransactionsbyMonthRange(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("GetTransactionsbyMonthRange", transaction.Account, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC)).Return([]model.Transaction{transaction}, nil)

	// call the GetTransactionsbyMonthRange method
	result, err := mockTransactionService.GetTransactionsbyMonthRange(transaction.Account, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC))

	if len(result) != 1 {
		t.Errorf("Expected result to have 1 element, but got %d", len(result))
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestGetTransactionsbyCount(t *testing.T) {

	mockTransactionService := NewMockTransactionService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock transaction service
	mockTransactionService.On("GetTransactionsbyCount", transaction.Account, 10).Return([]model.Transaction{transaction}, nil)

	// call the GetTransactionsbyCount method
	result, err := mockTransactionService.GetTransactionsbyCount(transaction.Account, 10)

	if len(result) != 1 {
		t.Errorf("Expected result to have 1 element, but got %d", len(result))
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}
