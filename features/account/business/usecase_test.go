package business

// import (
// 	"errors"
// 	"movie-api/features/account"
// 	"movie-api/features/account/mocks"
// 	"os"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// var (
// 	mockData        mocks.Data
// 	accountBusiness account.Business
// 	accountDatas    []account.AccountCore
// 	accountData     account.AccountCore
// 	accountLogin    account.AccountCore
// 	// accountUpdate account.AccountCore
// )

// func TestMain(m testing.M) {
// 	accountBusiness = NewBusinessAccount(&mockData)

// 	accountDatas = []account.AccountCore{
// 		{
// 			ID:       1,
// 			Username: "pertama",
// 			Password: "admin",
// 			Email:    "rasyid.id3@gmail.com",
// 		},
// 	}

// 	accountData = account.AccountCore{
// 		Username: "pertama",
// 		Password: "admin",
// 		Email:    "rasyid.id3@gmail.com",
// 	}

// 	accountLogin = account.AccountCore{
// 		Username: "pertama",
// 		Password: "admin",
// 	}

// 	os.Exit(m.Run)
// }

// func TestGetAccount(t *testing.T) {
// 	t.Run("validate get accounts", func(t *testing.T) {
// 		mockData.On("SelectAccount", mock.AnythingOfType("account.AccountCore")).Return(accountDatas, nil).Once()
// 		resp, err := accountBusiness.GetAccount(account.AccountCore{})
// 		assert.Nil(t, err)
// 		assert.Equal(t, len(resp), 1)
// 	})

// 	t.Run("error get accounts", func(t *testing.T) {
// 		mockData.On("SelectAccount", mock.AnythingOfType("account.AccountCore")).Return(nil, errors.New("error"))
// 		resp, err := accountBusiness.GetAccount(account.AccountCore{})
// 		assert.NotNil(t, err)
// 		assert.Nil(t, resp)
// 	})
// }
