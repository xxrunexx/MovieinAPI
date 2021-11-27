package data

import (
	"movie-api/features/movie"

	"gorm.io/gorm"
)

type Movie struct {
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

func toMovieRecord(movie movie.MovieCore) Movie {
	return Movie{
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

func toMovieCore(mv Movie) movie.MovieCore {
	return movie.MovieCore{
		Title:            mv.Title,
		Genre:            mv.Genre,
		ImdbID:           mv.ImdbID,
		OriginalLanguage: mv.OriginalLanguage,
		Price:            mv.Price,
		Popularity:       mv.Popularity,
		Status:           mv.Status,
		Overview:         mv.Overview,
		Vote_average:     mv.Vote_average,
		Vote_count:       mv.Vote_count,
	}
}

func toMovieCoreList(mvList []Movie) []movie.MovieCore {
	convMv := []movie.MovieCore{}

	for _, movie := range mvList {
		convMv = append(convMv, toMovieCore(movie))
	}
	return convMv
}
