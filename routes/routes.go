package routes

import (
	"movie-api/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	// Initiate Echo
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Account
	e.POST("/account/register", presenter.AccountPresentation.CreateAccountHandler)
	e.POST("/account/login", presenter.AccountPresentation.LoginAccountHandler)
	e.GET("/account", presenter.AccountPresentation.GetAccountsHandler)
	// e.PUT("/account/:id", UpdateAccount)

	// Watchlist
	e.POST("/account/:id/watchlist", presenter.WatchlistPresentation.CreateWatchlist)
	// e.GET("/account", GetAllAccount)
	// e.GET("/account/:id", GetAccountByID)

	return e
}
