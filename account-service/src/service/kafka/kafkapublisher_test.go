package service

import (
	"testing"
	"time"

	"github.com/banking-app/account-service/src/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockKafkaService struct {
	t *testing.T
	mock.Mock
}

func (m *mockKafkaService) PublishTransaction(transaction *model.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func NewMockKafkaService(t *testing.T) *mockKafkaService {
	return &mockKafkaService{t: t}
}

func TestPublishTransaction(t *testing.T) {

	mockKafkaService := NewMockKafkaService(t)

	transaction := model.Transaction{
		ID:        uuid.New().String(),
		Account:   uuid.New().String(),
		Amount:    100,
		Type:      "credit",
		Timestamp: time.Now(),
	}

	// add the transaction to the mock kafka service
	mockKafkaService.On("PublishTransaction", &transaction).Return(nil)

	// call the PublishTransaction method
	err := mockKafkaService.PublishTransaction(&transaction)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}