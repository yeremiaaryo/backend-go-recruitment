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

func Test_movieApplication_FindMovieByKeyword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error do request", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, "batman", "1")
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
		resp, err := ma.FindMovieByKeyword("batman", "1")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success do request, failed unmarshal", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, "batman", "1")
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
		resp, err := ma.FindMovieByKeyword("batman", "1")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success do request, success unmarshal, failed convert pagination", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, "batman", "abc")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)
		cm.EXPECT().Do(request).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`{}`))}, nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByKeyword("batman", "abc")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success do request, success unmarshal, success convert pagination, failed convert total result", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, "batman", "1")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)
		cm.EXPECT().Do(request).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`{}`))}, nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByKeyword("batman", "1")
		assert.Nil(t, resp)
		assert.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)

		url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, "batman", "1")
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return
		}
		mr := movie.NewMockMovieRepository(ctrl)
		cm.EXPECT().Do(request).Return(&http.Response{Body: ioutil.NopCloser(bytes.NewBufferString(`{"Search":[{"Title":"Casey at the Bat","Year":"1946","imdbID":"tt0038399","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BMTU1MTc5MDg3OV5BMl5BanBnXkFtZTcwMDI1NTQzMQ@@._V1_SX300.jpg"},{"Title":"The Sea Bat","Year":"1930","imdbID":"tt0021345","Type":"movie","Poster":"https://m.media-amazon.com/images/M/MV5BM2JmYjU3ODYtZWFhOS00ZWFhLWIwMWQtMWZmYTgyM2MwNjUxXkEyXkFqcGdeQXVyNDAzOTcxOTE@._V1_SX300.jpg"}],"totalResults":"221","Response":"True"}`))}, nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		resp, err := ma.FindMovieByKeyword("batman", "1")
		expectation := &movie2.MovieList{
			Search: []movie2.Movie{
				{
					Title:  "Casey at the Bat",
					Year:   "1946",
					ImdbID: "tt0038399",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BMTU1MTc5MDg3OV5BMl5BanBnXkFtZTcwMDI1NTQzMQ@@._V1_SX300.jpg",
					Detail: nil,
				},
				{
					Title:  "The Sea Bat",
					Year:   "1930",
					ImdbID: "tt0021345",
					Type:   "movie",
					Poster: "https://m.media-amazon.com/images/M/MV5BM2JmYjU3ODYtZWFhOS00ZWFhLWIwMWQtMWZmYTgyM2MwNjUxXkEyXkFqcGdeQXVyNDAzOTcxOTE@._V1_SX300.jpg",
					Detail: nil,
				},
			},
			TotalResults: "221",
			Response:     "True",
			TotalPage:    111,
			Page:         1,
		}
		assert.Equal(t, expectation, resp)
		assert.Nil(t, err)
	})
}
