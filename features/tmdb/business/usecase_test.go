package business

import (
	"errors"
	"movie-api/features/tmdb"
	"movie-api/features/tmdb/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData     mocks.Data
	tmdbBusiness tmdb.Business
	tmdbDatas    []tmdb.TmdbAPICore
)

func TestMain(m *testing.M) {
	tmdbBusiness = NewBusinessTmdb(&mockData)

	tmdbDatas = []tmdb.TmdbAPICore{}
	os.Exit(m.Run())
}

func TestGetMovieByTitle(t *testing.T) {
	t.Run("validate Get movie by title", func(t *testing.T) {
		mockData.On("SelectMovieByTitle", mock.AnythingOfType("string")).Return(tmdb.TmdbAPICore{}, nil).Once()
		resp, err := tmdbBusiness.GetMovieByTitle("Hitman")
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error", func(t *testing.T) {
		mockData.On("SelectMovieByTitle", mock.AnythingOfType("string")).Return(tmdb.TmdbAPICore{}, errors.New("error")).Once()
		resp, err := tmdbBusiness.GetMovieByTitle("Hitman")
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestGetMoviePopular(t *testing.T) {
	t.Run("validate Get popular movie", func(t *testing.T) {
		mockData.On("SelectMoviePopular").Return(tmdb.TmdbAPICore{}, nil).Once()
		resp, err := tmdbBusiness.GetMoviePopular()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error", func(t *testing.T) {
		mockData.On("SelectMoviePopular").Return(tmdb.TmdbAPICore{}, errors.New("error")).Once()
		resp, err := tmdbBusiness.GetMoviePopular()
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestGetMovieOnGoing(t *testing.T) {
	t.Run("validate Get on-going movie", func(t *testing.T) {
		mockData.On("SelectMovieOnGoing").Return(tmdb.TmdbAPICore{}, nil).Once()
		resp, err := tmdbBusiness.GetMovieOnGoing()
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error", func(t *testing.T) {
		mockData.On("SelectMovieOnGoing").Return(tmdb.TmdbAPICore{}, errors.New("error")).Once()
		resp, err := tmdbBusiness.GetMovieOnGoing()
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}
