package movie

import (
	movieApplication "backend-go-recruitment/application/movie"
	"backend-go-recruitment/common/movie"
)

type MovieInterface interface {
	FindMovieByKeyword(keyword, pagination string) (*movie.MovieList, error)
}

type movieInterface struct {
	movieApplication movieApplication.MovieApplication
}

func NewMovieInterface(ma movieApplication.MovieApplication) MovieInterface {
	return &movieInterface{
		movieApplication: ma,
	}
}
