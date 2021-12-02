package business

import (
	"errors"
	"movie-api/features/account"
	"movie-api/features/account/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData        mocks.Data
	accountBusiness account.Business
	accountDatas    []account.AccountCore
	accountData     account.AccountCore
	accountLogin    account.AccountCore
)

func TestMain(m *testing.M) {
	accountBusiness = NewBusinessAccount(&mockData)

	accountDatas = []account.AccountCore{
		{
			ID:       1,
			Username: "pertama",
			Password: "admin",
			Email:    "rasyid.id3@gmail.com",
		},
	}

	accountData = account.AccountCore{
		Username: "pertama",
		Password: "admin",
		Email:    "rasyid.id3@gmail.com",
	}

	accountLogin = account.AccountCore{
		Username: "pertama",
		Password: "admin",
	}

	os.Exit(m.Run())
}

func TestCreateAccount(t *testing.T) {
	t.Run("validate create account", func(t *testing.T) {
		mockData.On("InsertAccount", mock.AnythingOfType("account.AccountCore")).Return(nil).Once()
		err := accountBusiness.CreateAccount(account.AccountCore{})
		assert.Nil(t, err)
	})
	t.Run("error create account", func(t *testing.T) {
		mockData.On("InsertAccount", mock.AnythingOfType("account.AccountCore")).Return(errors.New("Error")).Once()
		err := accountBusiness.CreateAccount(account.AccountCore{})
		assert.NotNil(t, err)
	})
}
func TestGetAccount(t *testing.T) {
	t.Run("validate get accounts", func(t *testing.T) {
		mockData.On("SelectAccount", mock.AnythingOfType("account.AccountCore")).Return(accountDatas, nil).Once()
		resp, err := accountBusiness.GetAccount(account.AccountCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get accounts", func(t *testing.T) {
		mockData.On("SelectAccount", mock.AnythingOfType("account.AccountCore")).Return(nil, errors.New("error"))
		resp, err := accountBusiness.GetAccount(account.AccountCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestGetAccountByID(t *testing.T) {
	t.Run("validate get account by id", func(t *testing.T) {
		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(account.AccountCore{}, nil).Once()
		resp, err := accountBusiness.GetAccountByID(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("error get account by id", func(t *testing.T) {
		mockData.On("SelectAccountByID", mock.AnythingOfType("int")).Return(account.AccountCore{}, errors.New("error")).Once()
		resp, err := accountBusiness.GetAccountByID(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestLoginAccount(t *testing.T) {
	t.Run("validate login account", func(t *testing.T) {
		mockData.On("CheckAccount", mock.AnythingOfType("account.AccountCore")).Return(account.AccountCore{}, nil).Once()
		resp, err := accountBusiness.LoginAccount(account.AccountCore{})
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error login account", func(t *testing.T) {
		mockData.On("CheckAccount", mock.AnythingOfType("account.AccountCore")).Return(account.AccountCore{}, errors.New("error")).Once()
		resp, err := accountBusiness.LoginAccount(account.AccountCore{})
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}
