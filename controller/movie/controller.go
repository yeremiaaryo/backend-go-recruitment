package movie

import (
	"backend-go-recruitment/interfaces/movie"
)

type MovieController struct {
	MovieInterface movie.MovieInterface
}

func NewMovieController(mi movie.MovieInterface) *MovieController {
	return &MovieController{
		MovieInterface: mi,
	}
}
