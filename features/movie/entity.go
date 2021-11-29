package movie

type MovieCore struct {
	ID               int
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

// Untuk layer data / repo
type Data interface {
	SelectMovieByTitle(title string) (movie MovieCore, err error)
	SelectMoviePopular() (movie []MovieCore, err error)
	SelectMovieOnGoing() (movie []MovieCore, err error)
}

// Untuk layer business / service
type Business interface {
	GetMovieByTitle(title string) (movie MovieCore, err error)
	GetMoviePopular() (movie []MovieCore, err error)
	GetMovieOnGoing() (movie []MovieCore, err error)
}
