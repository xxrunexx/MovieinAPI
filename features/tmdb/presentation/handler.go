package presentation

import (
	"movie-api/features/tmdb"
	"movie-api/features/tmdb/presentation/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TmdbHandler struct {
	tmdbBusiness tmdb.Business
}

func NewHandlerTmdb(tmdbBusiness tmdb.Business) *TmdbHandler {
	return &TmdbHandler{tmdbBusiness}
}

func (tmdbHandler *TmdbHandler) GetMovieByTitleHandler(e echo.Context) error {
	title := e.QueryParam("title")

	movie, err := tmdbHandler.tmdbBusiness.GetMovieByTitle(title)

	if err != nil {
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    response.ToTmdbResponse(movie),
	})
}
