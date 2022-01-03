package movie

import (
	movie2 "backend-go-recruitment/common/movie"
	"backend-go-recruitment/common/movie/movie_grpc"
	"backend-go-recruitment/interfaces/movie"
	"context"
	"errors"
)

type MovieServer struct {
	movieInterface movie.MovieInterface
}

func NewMovieServer(mi movie.MovieInterface) movie_grpc.MovieServer {
	return MovieServer{
		movieInterface: mi,
	}
}

func (ms MovieServer) Search(ctx context.Context, param *movie_grpc.MovieRequest) (*movie_grpc.MovieList, error) {
	movie, err := ms.movieInterface.FindMovieByKeyword(param.Searchword, param.Pagination)
	if err != nil {
		return new(movie_grpc.MovieList), errors.New("failed")
	}

	return buildMovieSearchResponseGRPC(movie), err
}

func buildMovieSearchResponseGRPC(movie *movie2.MovieList) *movie_grpc.MovieList {
	movieList := make([]*movie_grpc.Search, 0)
	for _, v := range movie.Search {
		detail := new(movie_grpc.MovieDetail)
		if v.Detail != nil {
			ratings := make([]*movie_grpc.MovieRatings, 0)
			for _, v2 := range v.Detail.Ratings {
				rate := &movie_grpc.MovieRatings{
					Source: v2.Source,
					Value:  v2.Value,
				}
				ratings = append(ratings, rate)
			}
			detail = &movie_grpc.MovieDetail{
				Rated:      v.Detail.Rated,
				Released:   v.Detail.Released,
				Runtime:    v.Detail.Runtime,
				Genre:      v.Detail.Genre,
				Director:   v.Detail.Director,
				Writer:     v.Detail.Writer,
				Actors:     v.Detail.Actors,
				Plot:       v.Detail.Plot,
				Language:   v.Detail.Language,
				Country:    v.Detail.Country,
				Awards:     v.Detail.Awards,
				Poster:     v.Detail.Poster,
				Ratings:    ratings,
				Metascore:  v.Detail.Metascore,
				ImdbRating: v.Detail.ImdbRating,
				ImdbVotes:  v.Detail.ImdbVotes,
				ImdbID:     v.Detail.ImdbID,
				Type:       v.Detail.Type,
				DVD:        v.Detail.Dvd,
				BoxOffice:  v.Detail.BoxOffice,
				Production: v.Detail.Production,
				Website:    v.Detail.Website,
				Response:   v.Detail.Response,
			}
		}
		search := &movie_grpc.Search{
			Title:       v.Title,
			Year:        v.Year,
			ImdbID:      v.ImdbID,
			Type:        v.Type,
			Poster:      v.Poster,
			MovieDetail: detail,
		}
		movieList = append(movieList, search)
	}
	return &movie_grpc.MovieList{
		Search:       movieList,
		TotalResults: movie.TotalResults,
		Response:     movie.Response,
		TotalPage:    int64(movie.TotalPage),
		Page:         int64(movie.Page),
	}
}
