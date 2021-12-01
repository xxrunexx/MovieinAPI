package request

import "movie-api/features/tmdb"

type ReqTmdb struct {
	Title string `json:"title"`
}

func (movie *ReqTmdb) ToMovieApiCore() tmdb.TmdbAPICore {
	return tmdb.TmdbAPICore{
		Title: movie.Title,
	}
}
