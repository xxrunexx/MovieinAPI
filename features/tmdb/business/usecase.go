package business

import (
	"movie-api/features/tmdb"
)

type MovieAPIBusiness struct {
	movieAPIData tmdb.Data
}

func NewBusinessTmdb(tmdbData tmdb.Data) tmdb.Business {
	return &MovieAPIBusiness{tmdbData}
}

func (mvBusiness *MovieAPIBusiness) GetMovieByTitle(title string) (tmdb.TmdbAPICore, error) {
	movieData, err := mvBusiness.movieAPIData.SelectMovieByTitle(title)

	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}
	return movieData, nil
}
