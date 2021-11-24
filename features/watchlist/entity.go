package watchlist

import "time"

type WatchlistCore struct {
	ID        int
	AccountID int
	MovieID   int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Untuk layer data
type Data interface {
	InsertWatchlist(wlData WatchlistCore) error
}

// Untuk layer business
type Business interface {
	CreateWatchlist(wlData WatchlistCore) error
}
