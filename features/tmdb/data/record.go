package data

import (
	"movie-api/features/tmdb"
)

type MovieAPI struct {
	Page    int `json:"page"`
	Results []MovieData
}

type MovieData struct {
	ID               uint    `json:"id"`
	Title            string  `json:"title"`
	OriginalLanguage string  `json:"original_language"`
	Price            int     `json:"price"`
	Popularity       float32 `json:"popularity"`
	Overview         string  `json:"overview"`
	VoteAverage      float32 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

func (movie *MovieAPI) toMovieApiCore() tmdb.TmdbAPICore {
	return tmdb.TmdbAPICore{
		ID:               movie.Results[0].ID,
		Title:            movie.Results[0].Title,
		OriginalLanguage: movie.Results[0].OriginalLanguage,
		Price:            movie.Results[0].Price,
		Popularity:       movie.Results[0].Popularity,
		Overview:         movie.Results[0].Overview,
		VoteAverage:      movie.Results[0].VoteAverage,
		VoteCount:        movie.Results[0].VoteCount,
	}
}
