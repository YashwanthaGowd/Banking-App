package service

import (
	"testing"
	"time"

	"github.com/banking-app/account-service/src/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type mockAccountService struct {
	t *testing.T
	mock.Mock
}

func (m *mockAccountService) GetAccountbyId(accountId string) (*model.Account, error) {
	args := m.Called(accountId)
	return args.Get(0).(*model.Account), args.Error(1)
}

func (m *mockAccountService) CreateAccount(account *model.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *mockAccountService) UpdateAccount(account *model.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *mockAccountService) Deposit(accountID string, amount float64) error {
	args := m.Called(accountID, amount)
	return args.Error(0)
}

func (m *mockAccountService) Withdraw(accountID string, amount float64) error {
	args := m.Called(accountID, amount)
	return args.Error(0)
}

func NewMockAccountService(t *testing.T) *mockAccountService {
	return &mockAccountService{t: t}
}

func TestGetAccountbyId(t *testing.T) {

	mockAccountService := NewMockAccountService(t)


	account := model.Account{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Type:      "personal",
		Balance:   100,
		Status:    "active",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add the account to the mock account service
	mockAccountService.On("GetAccountbyId", account.ID).Return(&account, nil)

	// call the GetAccountbyId method
	result, err := mockAccountService.GetAccountbyId(account.ID)

	if result.ID != account.ID {
		t.Errorf("Expected result.ID to be %s, but got %s", account.ID, result.ID)
	}

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestCreateAccount(t *testing.T) {

	mockAccountService := NewMockAccountService(t)

	account := model.Account{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Type:      "personal",
		Balance:   100,
		Status:    "active",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add the account to the mock account service
	mockAccountService.On("CreateAccount", &account).Return(nil)

	// call the CreateAccount method
	err := mockAccountService.CreateAccount(&account)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestUpdateAccount(t *testing.T) {

	mockAccountService := NewMockAccountService(t)

	account := model.Account{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Type:      "personal",
		Balance:   100,
		Status:    "active",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add the account to the mock account service
	mockAccountService.On("UpdateAccount", &account).Return(nil)

	// call the UpdateAccount method
	err := mockAccountService.UpdateAccount(&account)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestDeposit(t *testing.T) {

	mockAccountService := NewMockAccountService(t)

	account := model.Account{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Type:      "personal",
		Balance:   100,
		Status:    "active",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add the account to the mock account service
	mockAccountService.On("Deposit", account.ID, float64(100)).Return(nil)

	// call the Deposit method
	err := mockAccountService.Deposit(account.ID, 100)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}

func TestWithdraw(t *testing.T) {

	mockAccountService := NewMockAccountService(t)

	account := model.Account{
		ID:        uuid.New().String(),
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Type:      "personal",
		Balance:   100,
		Status:    "active",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// add the account to the mock account service
	mockAccountService.On("Withdraw", account.ID, float64(100)).Return(nil)

	// call the Withdraw method
	err := mockAccountService.Withdraw(account.ID, 100)

	if err != nil {
		t.Errorf("Expected error to be nil, but got %v", err)
	}
}
