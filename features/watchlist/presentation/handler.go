package presentation

import (
	"movie-api/features/watchlist"
	"movie-api/features/watchlist/presentation/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WatchlistHandler struct {
	watchlistBusiness watchlist.Business
}

func NewHandlerWatchlist(watchlistBusiness watchlist.Business) *WatchlistHandler {
	return &WatchlistHandler{watchlistBusiness}
}

func (wlHandler *WatchlistHandler) CreateWatchlist(c echo.Context) error {
	newWatchlist := request.ReqWatchlist{}
	id, _ := strconv.Atoi(c.Param("id"))
	newWatchlist.MovieID = id
	if err := c.Bind(&newWatchlist); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	if err := wlHandler.watchlistBusiness.CreateWatchlist(request.ToWatchlistCore(newWatchlist)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Success",
		"data":    newWatchlist,
	})
}
