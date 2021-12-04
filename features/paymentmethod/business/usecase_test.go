package business

import (
	"errors"
	"movie-api/features/paymentmethod"
	"movie-api/features/paymentmethod/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData    mocks.Data
	pmBussiness paymentmethod.Business
	// pmDatas     []paymentmethod.PaymentmethodtCore
	// pmData paymentmethod.PaymentmethodCore
)

func TestMain(m *testing.M) {
	pmBussiness = NewBusinessPaymentmethod(&mockData)

	// pmData = []paymentmethod.PaymentmethodtCore{}

	os.Exit(m.Run())
}

func TestCreatePaymentmethod(t *testing.T) {
	t.Run("validate create payment method", func(t *testing.T) {
		mockData.On("InsertPaymentmethod", mock.AnythingOfType("paymentmethod.PaymentmethodCore")).Return(nil).Once()
		err := pmBussiness.CreatePaymentmethod(paymentmethod.PaymentmethodCore{})
		assert.Nil(t, err)
	})

	t.Run("error create payment method", func(t *testing.T) {
		mockData.On("InsertPaymentmethod", mock.AnythingOfType("paymentmethod.PaymentmethodCore")).Return(errors.New("error")).Once()
		err := pmBussiness.CreatePaymentmethod(paymentmethod.PaymentmethodCore{})
		assert.NotNil(t, err)
	})
}

func TestGetPaymentmethod(t *testing.T) {
	t.Run("validate get payment method", func(t *testing.T) {
		mockData.On("SelectPaymentmethod", mock.AnythingOfType("int")).Return(paymentmethod.PaymentmethodCore{}, nil).Once()
		resp, err := pmBussiness.GetPaymentmethod(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error get payment method", func(t *testing.T) {
		mockData.On("SelectPaymentmethod", mock.AnythingOfType("int")).Return(paymentmethod.PaymentmethodCore{}, errors.New("error")).Once()
		resp, err := pmBussiness.GetPaymentmethod(3)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestDeletePaymentmethod(t *testing.T) {
	t.Run("validate delete payment method", func(t *testing.T) {
		mockData.On("DeletePaymentmethod", mock.AnythingOfType("int")).Return(nil).Once()
		err := pmBussiness.DeletePaymentmethod(3)
		assert.Nil(t, err)
	})

	t.Run("error delete payment method", func(t *testing.T) {
		mockData.On("DeletePaymentmethod", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := pmBussiness.DeletePaymentmethod(3)
		assert.NotNil(t, err)
	})
}
