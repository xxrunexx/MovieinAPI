package routes

import (
	"movie-api/config"
	"movie-api/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	// Initiate Echo & JWT
	e := echo.New()
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// Account
	e.POST("/account/register", presenter.AccountPresentation.CreateAccountHandler)
	e.POST("/account/login", presenter.AccountPresentation.LoginAccountHandler)
	jwt.GET("/account", presenter.AccountPresentation.GetAccountsHandler)
	jwt.GET("/account/:id", presenter.AccountPresentation.GetAccountByIDHandler)
	// e.PUT("/account/:id", UpdateAccount)

	// Watchlist
	e.POST("/account/:id/watchlist", presenter.WatchlistPresentation.CreateWatchlistHandler)
	e.GET("/account/:account_id", presenter.WatchlistPresentation.GetWatchlistHandler)
	// e.GET("/account/:id", GetAccountByID)

	// Movie
	e.GET("/movie/:title", presenter.TmdbPresentation.GetMovieByTitleHandler)
	jwt.GET("/movie/popular", presenter.TmdbPresentation.GetMoviePopularHandler)
	jwt.GET("/movie/ongoing", presenter.TmdbPresentation.GetMovieOnGoingHandler)

	// Transaction

	return e
}
