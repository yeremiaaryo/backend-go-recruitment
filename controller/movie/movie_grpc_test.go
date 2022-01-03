package movie

import (
	movie2 "backend-go-recruitment/common/movie"
	"backend-go-recruitment/common/movie/movie_grpc"
	"backend-go-recruitment/mocks/interfaces/movie"
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMovieServer_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When failed", func(t *testing.T) {
		mi := movie.NewMockMovieInterface(ctrl)
		mi.EXPECT().FindMovieByKeyword("batman", "1").Return(nil, errors.New("failed"))

		ms := MovieServer{movieInterface: mi}

		resp, err := ms.Search(context.Background(), &movie_grpc.MovieRequest{
			Searchword: "batman",
			Pagination: "1",
		})
		expectation := new(movie_grpc.MovieList)
		assert.Equal(t, err, errors.New("failed"))
		assert.Equal(t, expectation, resp)
	})

	t.Run("When success", func(t *testing.T) {
		mi := movie.NewMockMovieInterface(ctrl)
		mi.EXPECT().FindMovieByKeyword("batman", "1").Return(&movie2.MovieList{
			Search: []movie2.Movie{
				{
					Title:  "Batman Begins",
					Year:   "2005",
					ImdbID: "tt0372784",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
					Detail: &movie2.MovieDetail{
						Rated:    "PG-13",
						Released: "25 Mar 2016",
						Runtime:  "152 min",
						Genre:    "Action, Adventure, Sci-Fi",
						Director: "Zack Snyder",
						Writer:   "Chris Terrio, David S. Goyer, Bob Kane",
						Actors:   "Ben Affleck, Henry Cavill, Amy Adams",
						Plot:     "Fearing that the actions of Superman are left unchecked, Batman takes on the Man of Steel, while the world wrestles with what kind of a hero it really needs.",
						Language: "English",
						Country:  "United States",
						Awards:   "14 wins & 33 nominations",
						Poster:   "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
						Ratings: []movie2.MovieRatings{
							{
								Source: "Internet Movie Database",
								Value:  "8.2/10",
							},
						},
						Metascore:  "44",
						ImdbRating: "6.4",
						ImdbVotes:  "670,863",
						ImdbID:     "tt2975590",
						Type:       "movie",
						Dvd:        "19 Jul 2016",
						BoxOffice:  "$330,360,194",
						Production: "N/A",
						Website:    "N/A",
						Response:   "True",
					},
				},
			},
			TotalResults: "490",
			Response:     "True",
			TotalPage:    49,
			Page:         1,
		}, nil)

		ms := MovieServer{movieInterface: mi}

		resp, err := ms.Search(context.Background(), &movie_grpc.MovieRequest{
			Searchword: "batman",
			Pagination: "1",
		})
		expectation := &movie_grpc.MovieList{
			Search: []*movie_grpc.Search{
				{
					Title:  "Batman Begins",
					Year:   "2005",
					ImdbID: "tt0372784",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
					MovieDetail: &movie_grpc.MovieDetail{
						Rated:    "PG-13",
						Released: "25 Mar 2016",
						Runtime:  "152 min",
						Genre:    "Action, Adventure, Sci-Fi",
						Director: "Zack Snyder",
						Writer:   "Chris Terrio, David S. Goyer, Bob Kane",
						Actors:   "Ben Affleck, Henry Cavill, Amy Adams",
						Plot:     "Fearing that the actions of Superman are left unchecked, Batman takes on the Man of Steel, while the world wrestles with what kind of a hero it really needs.",
						Language: "English",
						Country:  "United States",
						Awards:   "14 wins & 33 nominations",
						Poster:   "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
						Ratings: []*movie_grpc.MovieRatings{
							{
								Source: "Internet Movie Database",
								Value:  "8.2/10",
							},
						},
						Metascore:  "44",
						ImdbRating: "6.4",
						ImdbVotes:  "670,863",
						ImdbID:     "tt2975590",
						Type:       "movie",
						DVD:        "19 Jul 2016",
						BoxOffice:  "$330,360,194",
						Production: "N/A",
						Website:    "N/A",
						Response:   "True",
					},
				},
			},
			TotalResults: "490",
			Response:     "True",
			TotalPage:    49,
			Page:         1,
		}
		assert.Nil(t, err)
		assert.Equal(t, expectation, resp)
	})
}

func TestNewMovieServer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mi := movie.NewMockMovieInterface(ctrl)
	interFace := NewMovieServer(mi)
	assert.IsType(t, MovieServer{}, interFace)

}
