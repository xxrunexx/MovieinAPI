package data

import (
	"encoding/json"
	"io/ioutil"
	"movie-api/features/tmdb"
	"net/http"
	"net/url"
	"strings"
)

type Tmdb struct {
	httpClient http.Client
}

func NewData() tmdb.Data {
	return &Tmdb{
		httpClient: http.Client{},
	}
}

// Third Party
func (api *Tmdb) SelectMovieByTitle(title string) (tmdb.TmdbAPICore, error) {
	var movie MovieAPI
	processedData := strings.ToLower(title)
	endPoint := "https://api.themoviedb.org/3/search/movie"
	req, err := http.NewRequest("GET", endPoint+url.QueryEscape(processedData), nil)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}

	defer resp.Body.Close()
	bodybytes, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(bodybytes, &movie)

	return movie.toMovieApiCore(), nil

	// Archieve
	// defer resp.Body.Close()
	// bodybytes, _ := io.ReadAll(resp.Body)
	// var movietitle []MovieAPI
	// _ = json.Unmarshal(bodybytes, &movietitle)

	// if err != nil {
	// 	return tmdb.TmdbAPICore{}, err
	// }
	// // fmt.Println("Data : " + movietitle[])
	// return movietitle[0].toMovieApiCore(), nil
}
