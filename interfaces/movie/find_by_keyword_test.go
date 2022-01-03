package movie

import (
	movie2 "backend-go-recruitment/common/movie"
	"backend-go-recruitment/mocks/application/movie"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_movieInterface_FindMovieByKeyword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error find by keyword", func(t *testing.T) {
		ma := movie.NewMockMovieApplication(ctrl)
		ma.EXPECT().FindMovieByKeyword("batman", "1").Return(nil, errors.New("failed"))

		mi := movieInterface{movieApplication: ma}
		res, err := mi.FindMovieByKeyword("batman", "1")
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("When error find by keyword", func(t *testing.T) {
		ma := movie.NewMockMovieApplication(ctrl)
		ma.EXPECT().FindMovieByKeyword("batman", "1").Return(nil, nil)

		mi := movieInterface{movieApplication: ma}
		res, err := mi.FindMovieByKeyword("batman", "1")
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("When error find by imdb id, error insert logs, return success (not breaking)", func(t *testing.T) {
		ma := movie.NewMockMovieApplication(ctrl)
		respFind := &movie2.MovieList{
			Search: []movie2.Movie{
				{
					Title:  "Batman Begins",
					Year:   "2005",
					ImdbID: "tt0372784",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
				},
			},
			TotalResults: "490",
			Response:     "True",
			TotalPage:    49,
			Page:         1,
		}
		ma.EXPECT().FindMovieByKeyword("batman", "1").Return(respFind, nil)
		for _, v := range respFind.Search {
			ma.EXPECT().FindMovieByImdbID(v.ImdbID).Return(nil, errors.New("failed"))
		}
		ma.EXPECT().InsertLogs("batman", "1").Return(errors.New("failed"))

		mi := movieInterface{movieApplication: ma}
		res, err := mi.FindMovieByKeyword("batman", "1")
		expectation := &movie2.MovieList{
			Search: []movie2.Movie{
				{
					Title:  "Batman Begins",
					Year:   "2005",
					ImdbID: "tt0372784",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
				},
			},
			TotalResults: "490",
			Response:     "True",
			TotalPage:    49,
			Page:         1,
		}
		assert.Equal(t, expectation, res)
		assert.Nil(t, err)
	})

	t.Run("When all is success", func(t *testing.T) {
		ma := movie.NewMockMovieApplication(ctrl)
		respFind := &movie2.MovieList{
			Search: []movie2.Movie{
				{
					Title:  "Batman Begins",
					Year:   "2005",
					ImdbID: "tt0372784",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BOTY4YjI2N2MtYmFlMC00ZjcyLTg3YjEtMDQyM2ZjYzQ5YWFkXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
				},
			},
			TotalResults: "490",
			Response:     "True",
			TotalPage:    49,
			Page:         1,
		}
		detail := &movie2.MovieDetail{
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
		}
		ma.EXPECT().FindMovieByKeyword("batman", "1").Return(respFind, nil)
		for _, v := range respFind.Search {
			ma.EXPECT().FindMovieByImdbID(v.ImdbID).Return(detail, nil)
		}
		ma.EXPECT().InsertLogs("batman", "1").Return(nil)

		mi := movieInterface{movieApplication: ma}
		res, err := mi.FindMovieByKeyword("batman", "1")
		expectation := &movie2.MovieList{
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
		}
		assert.Equal(t, expectation, res)
		assert.Nil(t, err)
	})
}
