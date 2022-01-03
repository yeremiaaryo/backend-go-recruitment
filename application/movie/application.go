package movie

import (
	"backend-go-recruitment/common"
	"backend-go-recruitment/common/movie"
	movie2 "backend-go-recruitment/repository/movie"
	"net/http"
)

type MovieApplication interface {
	FindMovieByKeyword(keyword, pagination string) (*movie.MovieList, error)
	FindMovieByImdbID(imdbID string) (*movie.MovieDetail, error)
	InsertLogs(keyword, pagination string) error
}

type movieApplication struct {
	client          common.HttpClient
	movieRepository movie2.MovieRepository
}

func NewMovieApplication(mr movie2.MovieRepository) MovieApplication {
	client := new(http.Client)
	return &movieApplication{
		client:          client,
		movieRepository: mr,
	}
}
