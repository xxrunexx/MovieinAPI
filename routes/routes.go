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

	// Account
	e.POST("/account", presenter.AccountPresentation.CreateAccount)
	// e.PUT("/account/:id", UpdateAccount)

	// Watchlist
	e.POST("/account/:id/watchlist", presenter.WatchlistPresentation.CreateWatchlist)
	// e.GET("/account", GetAllAccount)
	// e.GET("/account/:id", GetAccountByID)

	return e
}