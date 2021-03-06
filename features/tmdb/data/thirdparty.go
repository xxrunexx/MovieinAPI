package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	fmt.Println("Isi processedData = " + processedData)
	params := url.Values{}
	params.Add("api_key", "c41ba5597d4df58bf5a33ea7151abd7a")
	params.Add("query", processedData)
	endPoint := "https://api.themoviedb.org/3/search/movie?"
	fmt.Println("Isi params : " + params.Encode())
	req, err := http.NewRequest("GET", endPoint+params.Encode(), nil)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}
	fmt.Println("Isi url string : " + req.URL.String())
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}

	defer resp.Body.Close()
	bodybytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(bodybytes))
	json.Unmarshal(bodybytes, &movie)
	return movie.toMovieApiCore(), nil
}

func (api *Tmdb) SelectMoviePopular() (tmdb.TmdbAPICore, error) {
	var movie MovieAPI
	params := url.Values{}
	params.Add("api_key", "c41ba5597d4df58bf5a33ea7151abd7a")
	endPoint := "https://api.themoviedb.org/3/movie/popular?"
	req, err := http.NewRequest("GET", endPoint+params.Encode(), nil)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}

	defer resp.Body.Close()
	bodybytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(bodybytes))
	json.Unmarshal(bodybytes, &movie)
	return movie.toMovieApiCore(), nil
}

func (api *Tmdb) SelectMovieOnGoing() (tmdb.TmdbAPICore, error) {
	var movie MovieAPI
	params := url.Values{}
	params.Add("api_key", "c41ba5597d4df58bf5a33ea7151abd7a")
	endPoint := "https://api.themoviedb.org/3/movie/now_playing?"
	req, err := http.NewRequest("GET", endPoint+params.Encode(), nil)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return tmdb.TmdbAPICore{}, err
	}

	defer resp.Body.Close()
	bodybytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(bodybytes))
	json.Unmarshal(bodybytes, &movie)
	return movie.toMovieApiCore(), nil
}
