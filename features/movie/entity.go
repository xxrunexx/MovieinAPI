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
	SelectMovie(mvData MovieCore) ([]MovieCore, error)
}

// Untuk layer business / service
type Business interface {
	GetMovie(mvData MovieCore) ([]MovieCore, error)
}
