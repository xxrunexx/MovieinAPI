package data

import (
	"movie-api/features/tmdb"

	"gorm.io/gorm"
)

// SelectMovieByTitle(title string) (movie MovieCore, err error)

type MovieAPI struct {
	gorm.Model
	Title            string
	Genre            string
	ImdbID           int
	OriginalLanguage string
	Price            int
	Popularity       float32
	Status           string
	Overview         string
	Vote_average     float32
	Vote_count       int
}

func (movie *MovieAPI) toMovieApiCore() tmdb.TmdbAPICore {
	return tmdb.TmdbAPICore{
		Title:            movie.Title,
		Genre:            movie.Genre,
		ImdbID:           movie.ImdbID,
		OriginalLanguage: movie.OriginalLanguage,
		Price:            movie.Price,
		Popularity:       movie.Popularity,
		Status:           movie.Status,
		Overview:         movie.Overview,
		Vote_average:     movie.Vote_average,
		Vote_count:       movie.Vote_count,
	}
}
