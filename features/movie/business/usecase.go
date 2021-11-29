package business

// type MovieBusiness struct {
// 	movieData movie.Data
// }

// func NewBusinessMovie(movieData movie.Data) movie.Business {
// 	return &MovieBusiness{movieData}
// }

// func (mvBusiness *MovieBusiness) GetMovieByTitle(title string) (movie.MovieCore, error) {
// 	movieData, err := mvBusiness.movieData.SelectMovieByTitle(title)
// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}
// 	return movieData, nil
// }

// func (mvBusiness *MovieBusiness) GetMoviePopular() (movie.MovieCore, error) {
// 	movieDatas, err := mvBusiness.movieData.SelectMoviePopular()
// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}
// 	return movieDatas, nil
// }

// func (mvBusiness *MovieBusiness) GetMovieOnGoing() (movie.MovieCore, error) {
// 	movieDatas, err := mvBusiness.movieData.SelectMovieOnGoing()
// 	if err != nil {
// 		return movie.MovieCore{}, err
// 	}
// 	return movieDatas, nil
// }
