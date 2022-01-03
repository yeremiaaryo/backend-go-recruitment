package movie

import (
	movie "backend-go-recruitment/mocks/repository/movie"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMovieApplication(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := movie.NewMockMovieRepository(ctrl)
	ma := NewMovieApplication(mr)
	assert.IsType(t, &movieApplication{}, ma)
}
