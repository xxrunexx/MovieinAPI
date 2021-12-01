package response

import "movie-api/features/tmdb"

type RespTmdb struct {
	ID               uint    `json:"id"`
	Title            string  `json:"title"`
	ImdbID           int     `json:"imdb_id"`
	OriginalLanguage string  `json:"original_language"`
	Price            int     `json:"price"`
	Popularity       float32 `json:"popularity"`
	Status           string  `json:"status"`
	Overview         string  `json:"overview"`
	VoteAverage      float32 `json:"voteAverage"`
	VoteCount        int     `json:"voteCount"`
}

func ToTmdbResponse(tmdb tmdb.TmdbAPICore) RespTmdb {
	return RespTmdb{
		ID:               tmdb.ID,
		Title:            tmdb.Title,
		ImdbID:           tmdb.ImdbID,
		OriginalLanguage: tmdb.OriginalLanguage,
		Price:            tmdb.Price,
		Popularity:       tmdb.Popularity,
		Status:           tmdb.Status,
		Overview:         tmdb.Overview,
		VoteAverage:      tmdb.VoteAverage,
		VoteCount:        tmdb.VoteCount,
	}
}
