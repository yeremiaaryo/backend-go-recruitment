package movie

import (
	movie2 "backend-go-recruitment/common/movie"
	"backend-go-recruitment/mocks/common"
	"backend-go-recruitment/mocks/repository/movie"
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func Test_movieApplication_FindMovieByImdbID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error do request", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", apiKey, "imdbID")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)

		cm.EXPECT().Do(request).Return(nil, errors.New("failed"))

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByImdbID("imdbID")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success do request, failed unmarshal", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", apiKey, "imdbID")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)
		cm.EXPECT().Do(request).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewReader(nil))}, nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByImdbID("imdbID")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", apiKey, "imdbID")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)
		cm.EXPECT().Do(request).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`{"Title":"The Sea Bat","Year":"1930","Rated":"Passed","Released":"05 Jul 1930","Runtime":"58 min","Genre":"Action, Romance, Thriller","Director":"Lionel Barrymore, Wesley Ruggles","Writer":"Dorothy Yost, Bess Meredyth, John Howard Lawson","Actors":"Raquel Torres, Charles Bickford, Nils Asther","Plot":"The West Indies island of Portuga exists mainly for sponge diving. But the best area of collection is frequented by a very large manta ray. Nina loses her brother to the creature and is comforted by a newly arrived minister, who seem","Language":"English","Country":"United States","Awards":"N/A","Poster":"https://m.media-amazon.com/images/M/MV5BM2JmYjU3ODYtZWFhOS00ZWFhLWIwMWQtMWZmYTgyM2MwNjUxXkEyXkFqcGdeQXVyNDAzOTcxOTE@._V1_SX300.jpg","Ratings":[{"Source":"Internet Movie Database","Value":"5.3/10"}],"Metascore":"N/A","imdbRating":"5.3","imdbVotes":"272","imdbID":"tt0021345","Type":"movie","DVD":"N/A","BoxOffice":"N/A","Production":"N/A","Website":"N/A","Response":"True"}`))}, nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByImdbID("imdbID")
		expectation := &movie2.MovieDetail{
			Rated:    "Passed",
			Released: "05 Jul 1930",
			Runtime:  "58 min",
			Genre:    "Action, Romance, Thriller",
			Director: "Lionel Barrymore, Wesley Ruggles",
			Writer:   "Dorothy Yost, Bess Meredyth, John Howard Lawson",
			Actors:   "Raquel Torres, Charles Bickford, Nils Asther",
			Plot:     "The West Indies island of Portuga exists mainly for sponge diving. But the best area of collection is frequented by a very large manta ray. Nina loses her brother to the creature and is comforted by a newly arrived minister, who seem",
			Language: "English",
			Country:  "United States",
			Awards:   "N/A",
			Poster:   "https://m.media-amazon.com/images/M/MV5BM2JmYjU3ODYtZWFhOS00ZWFhLWIwMWQtMWZmYTgyM2MwNjUxXkEyXkFqcGdeQXVyNDAzOTcxOTE@._V1_SX300.jpg",
			Ratings: []movie2.MovieRatings{
				{
					Source: "Internet Movie Database",
					Value:  "5.3/10",
				},
			},
			Metascore:  "N/A",
			ImdbRating: "5.3",
			ImdbVotes:  "272",
			ImdbID:     "tt0021345",
			Type:       "movie",
			Dvd:        "N/A",
			BoxOffice:  "N/A",
			Production: "N/A",
			Website:    "N/A",
			Response:   "True",
		}
		assert.Equal(t, expectation, resp)
		assert.Nil(t, err)
	})
}
