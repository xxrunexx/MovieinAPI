// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	tmdb "movie-api/features/tmdb"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// SelectMovieByTitle provides a mock function with given fields: title
func (_m *Data) SelectMovieByTitle(title string) (tmdb.TmdbAPICore, error) {
	ret := _m.Called(title)

	var r0 tmdb.TmdbAPICore
	if rf, ok := ret.Get(0).(func(string) tmdb.TmdbAPICore); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(tmdb.TmdbAPICore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectMovieOnGoing provides a mock function with given fields:
func (_m *Data) SelectMovieOnGoing() (tmdb.TmdbAPICore, error) {
	ret := _m.Called()

	var r0 tmdb.TmdbAPICore
	if rf, ok := ret.Get(0).(func() tmdb.TmdbAPICore); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tmdb.TmdbAPICore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectMoviePopular provides a mock function with given fields:
func (_m *Data) SelectMoviePopular() (tmdb.TmdbAPICore, error) {
	ret := _m.Called()

	var r0 tmdb.TmdbAPICore
	if rf, ok := ret.Get(0).(func() tmdb.TmdbAPICore); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(tmdb.TmdbAPICore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}