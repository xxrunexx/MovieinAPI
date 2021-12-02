package presentation

import (
	"fmt"
	"movie-api/features/watchlist"
	"movie-api/features/watchlist/presentation/request"
	"movie-api/features/watchlist/presentation/response"
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

func (wlHandler *WatchlistHandler) CreateWatchlistHandler(e echo.Context) error {
	newWatchlist := request.ReqWatchlist{}

	if err := e.Bind(&newWatchlist); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := wlHandler.watchlistBusiness.CreateWatchlist(newWatchlist.ToWatchlistCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newWatchlist,
	})
}

func (wlHandler *WatchlistHandler) GetWatchlistHandler(e echo.Context) error {
	account_id, err := strconv.Atoi(e.Param("account_id"))
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	data, err := wlHandler.watchlistBusiness.GetWatchlist(account_id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    response.ToWatchlistResponseList(data),
	})
}

func (wlHandler *WatchlistHandler) DeleteWatchlistHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	fmt.Println("Isi id : ", id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	data, err := wlHandler.watchlistBusiness.DeleteWatchlist(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "data deleted",
		"data":    response.ToWatchlistResponse(data),
	})
}
