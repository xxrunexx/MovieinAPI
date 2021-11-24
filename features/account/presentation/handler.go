package presentation

import (
	// Import Echo
	"net/http"

	"github.com/labstack/echo/v4"

	// Import files
	"movie-api/features/account"
	"movie-api/features/account/presentation/request"
	"movie-api/features/account/presentation/response"
)

type AccountHandler struct {
	accountBusiness account.Business
}

func NewHandlerAccount(accountBusiness account.Business) *AccountHandler {
	return &AccountHandler{accountBusiness}
}

func (accHandler *AccountHandler) CreateAccountHandler(e echo.Context) error {
	newAccount := request.ReqAccount{}

	if err := e.Bind(&newAccount); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := accHandler.accountBusiness.CreateAccount(newAccount.ToAccountCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newAccount,
	})
}

func (accHandler *AccountHandler) GetAccountsHandler(e echo.Context) error {
	data, err := accHandler.accountBusiness.GetAccount(account.AccountCore{})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToAccountResponseList(data),
	})
}
