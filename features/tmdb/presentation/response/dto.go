package response

import "movie-api/features/tmdb"

type RespTmdb struct {
	ID               uint    `json:"id"`
	Title            string  `json:"title"`
	Genre            string  `json:"genre"`
	ImdbID           int     `json:"imdb_id"`
	OriginalLanguage string  `json:"original_language"`
	Price            int     `json:"price"`
	Popularity       float32 `json:"popularity"`
	Status           string  `json:"status"`
	Overview         string  `json:"overview"`
	Vote_average     float32 `json:"vote_average"`
	Vote_count       int     `json:"vote_count"`
}

func ToTmdbResponse(tmdb tmdb.TmdbAPICore) RespTmdb {
	return RespTmdb{
		ID:               tmdb.ID,
		Title:            tmdb.Title,
		Genre:            tmdb.Genre,
		ImdbID:           tmdb.ImdbID,
		OriginalLanguage: tmdb.OriginalLanguage,
		Price:            tmdb.Price,
		Popularity:       tmdb.Popularity,
		Status:           tmdb.Status,
		Overview:         tmdb.Overview,
		Vote_average:     tmdb.Vote_average,
		Vote_count:       tmdb.Vote_count,
	}
}
