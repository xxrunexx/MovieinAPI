package tmdb

type TmdbAPICore struct {
	ID               uint
	Title            string
	Genre            string
	ImdbID           int
	OriginalLanguage string
	Price            int
	Popularity       float32
	Status           string
	Overview         string
	VoteAverage      float32
	VoteCount        int
}

type Data interface {
	SelectMovieByTitle(title string) (movie TmdbAPICore, err error)
	SelectMoviePopular() (movie TmdbAPICore, err error)
	SelectMovieOnGoing() (movie TmdbAPICore, err error)
}

// Untuk layer business / service
type Business interface {
	GetMovieByTitle(title string) (movie TmdbAPICore, err error)
	GetMoviePopular() (movie TmdbAPICore, err error)
	GetMovieOnGoing() (movie TmdbAPICore, err error)
}
