package business

import (
	"errors"
	"movie-api/features/watchlist"
	"movie-api/features/watchlist/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData           mocks.Data
	watchlistBussiness watchlist.Business
	// watchlistDatas     []watchlist.WatchlistCore
	// watchlistData      watchlist.WatchlistCore
)

func TestMain(m *testing.M) {
	watchlistBussiness = NewBusinessWatchlist(&mockData)

	// watchlistData = []watchlist.WatchlistCore{}

	os.Exit(m.Run())
}

func TestCreateWatchlist(t *testing.T) {
	t.Run("validate create watchlist", func(t *testing.T) {
		mockData.On("InsertWatchlist", mock.AnythingOfType("watchlist.WatchlistCore")).Return(nil).Once()
		err := watchlistBussiness.CreateWatchlist(watchlist.WatchlistCore{})
		assert.Nil(t, err)
	})

	t.Run("error create watchlist", func(t *testing.T) {
		mockData.On("InsertWatchlist", mock.AnythingOfType("watchlist.WatchlistCore")).Return(errors.New("error")).Once()
		err := watchlistBussiness.CreateWatchlist(watchlist.WatchlistCore{})
		assert.NotNil(t, err)
	})
}
func TestGetWatchlist(t *testing.T) {
	t.Run("validate get watchlist", func(t *testing.T) {
		mockData.On("SelectWatchlist", mock.AnythingOfType("int")).Return([]watchlist.WatchlistCore{}, nil).Once()
		resp, err := watchlistBussiness.GetWatchlist(3)
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 0)
	})

	t.Run("error get watchlist", func(t *testing.T) {
		mockData.On("SelectWatchlist", mock.AnythingOfType("int")).Return([]watchlist.WatchlistCore{}, errors.New("error")).Once()
		resp, err := watchlistBussiness.GetWatchlist(3)
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestDeleteWatchlist(t *testing.T) {
	t.Run("validate delete watchlist", func(t *testing.T) {
		mockData.On("DeleteWatchlist", mock.AnythingOfType("int")).Return(nil).Once()
		err := watchlistBussiness.DeleteWatchlist(3)
		assert.Nil(t, err)
	})

	t.Run("error delete watchlist", func(t *testing.T) {
		mockData.On("DeleteWatchlist", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := watchlistBussiness.DeleteWatchlist(3)
		assert.NotNil(t, err)
	})
}
