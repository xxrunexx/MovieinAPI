package data

import (
	"net/http"
)

type Tmdb struct {
	httpClient http.Client
}

// func NewData() movie.Data {
// 	return &Tmdb{
// 		httpClient: http.Client{},
// 	}
// }

// // Third Party
// func (tmdb *Tmdb) SelectMovieByID(title string) (movie.MovieCore, error) {
// 	processedData := strings.ToLower(title)
// 	endPoint := "https://api.themoviedb.org/3/search/movie"
// 	req, err := http.NewRequest("GET", endPoint+url.QueryEscape(processedData), nil)
// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}

// 	resp, err := tmdb.httpClient.Do(req)
// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}

// 	defer resp.Body.Close()
// 	bodybytes, _ := io.ReadAll(resp.Body)
// 	var movietitle []Movie
// 	_ = json.Unmarshal(bodybytes, &movietitle)

// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}
// 	// fmt.Println("Data : " + movietitle[])
// 	return movietitle[0].toMovieApiRecord(), nil
// }
