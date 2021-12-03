package business

import (
	"errors"
	"movie-api/features/transaction"
	"movie-api/features/transaction/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData            mocks.Data
	transactionBusiness transaction.Business
	transactionDatas    []transaction.TransactionCore
)

func TestMain(m *testing.M) {
	transactionBusiness = NewBusinessTransaction(&mockData)

	transactionDatas = []transaction.TransactionCore{}

	os.Exit(m.Run())
}

func TestCreateTransaction(t *testing.T) {
	t.Run("validate create transaction", func(t *testing.T) {
		mockData.On("InsertTransaction", mock.AnythingOfType("transaction.TransactionCore")).Return(nil).Once()
		err := transactionBusiness.CreateTransaction(transaction.TransactionCore{})
		assert.Nil(t, err)
	})

	t.Run("error create transaction", func(t *testing.T) {
		mockData.On("InsertTransaction", mock.AnythingOfType("transaction.TransactionCore")).Return(errors.New("error")).Once()
		err := transactionBusiness.CreateTransaction(transaction.TransactionCore{})
		assert.NotNil(t, err)
	})
}

func TestGetTransaction(t *testing.T) {
	t.Run("validate get transaction", func(t *testing.T) {
		mockData.On("SelectTransaction", mock.AnythingOfType("int")).Return([]transaction.TransactionCore{}, nil).Once()
		resp, err := transactionBusiness.GetTransaction(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error get transaction", func(t *testing.T) {
		mockData.On("SelectTransaction", mock.AnythingOfType("int")).Return([]transaction.TransactionCore{}, errors.New("error")).Once()
		resp, err := transactionBusiness.GetTransaction(3)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestDeleteTransaction(t *testing.T) {
	t.Run("validate delete transaction", func(t *testing.T) {
		mockData.On("DeleteTransaction", mock.AnythingOfType("int")).Return(nil).Once()
		err := transactionBusiness.DeleteTransaction(3)
		assert.Nil(t, err)
	})

	t.Run("error delete transaction", func(t *testing.T) {
		mockData.On("DeleteTransaction", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := transactionBusiness.DeleteTransaction(3)
		assert.NotNil(t, err)
	})
}
