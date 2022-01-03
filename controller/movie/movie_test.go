package movie

import (
	movie2 "backend-go-recruitment/common/movie"
	"backend-go-recruitment/mocks/interfaces/movie"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestMovieController_FindMovieByKeyword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("When error", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/?searchword=batman&pagination=3", nil)
		res := httptest.NewRecorder()

		mi := movie.NewMockMovieInterface(ctrl)
		mi.EXPECT().FindMovieByKeyword("batman", "3").Return(nil, errors.New("failed"))
		h := NewMovieController(mi)

		err := h.FindMovieByKeyword(e.NewContext(req, res))
		assert.Error(t, err)
	})

	t.Run("When success", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/?searchword=batman&pagination=3", nil)
		res := httptest.NewRecorder()

		mi := movie.NewMockMovieInterface(ctrl)
		mi.EXPECT().FindMovieByKeyword("batman", "3").Return(&movie2.MovieList{}, nil)
		h := NewMovieController(mi)

		err := h.FindMovieByKeyword(e.NewContext(req, res))
		assert.Nil(t, err)
	})
}
