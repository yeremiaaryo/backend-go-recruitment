package movie

import (
	movie2 "backend-go-recruitment/mocks/application/movie"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMovieInterface(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ma := movie2.NewMockMovieApplication(ctrl)
	interFace := NewMovieInterface(ma)
	assert.IsType(t, &movieInterface{}, interFace)
}
