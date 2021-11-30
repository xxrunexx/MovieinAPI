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
	Vote_average     float32
	Vote_count       int
}

type Data interface {
	SelectMovieByTitle(title string) (movie TmdbAPICore, err error)
	// SelectMoviePopular() (movie []MovieCore, err error)
	// SelectMovieOnGoing() (movie []MovieCore, err error)
}

// Untuk layer business / service
type Business interface {
	GetMovieByTitle(title string) (movie TmdbAPICore, err error)
	// 	GetMoviePopular() (movie []MovieCore, err error)
	// 	GetMovieOnGoing() (movie []MovieCore, err error)
}
