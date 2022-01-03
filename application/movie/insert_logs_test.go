package movie

import (
	"backend-go-recruitment/mocks/common"
	"backend-go-recruitment/mocks/repository/movie"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_movieApplication_InsertLogs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("when error", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)
		mr := movie.NewMockMovieRepository(ctrl)

		mr.EXPECT().Insert("batman", "1").Return(errors.New("failed"))

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		err := ma.InsertLogs("batman", "1")
		assert.Error(t, err)
	})

	t.Run("when error", func(t *testing.T) {
		cm := common.NewMockHttpClient(ctrl)
		mr := movie.NewMockMovieRepository(ctrl)

		mr.EXPECT().Insert("batman", "1").Return(nil)

		ma := movieApplication{
			client:          cm,
			movieRepository: mr,
		}
		err := ma.InsertLogs("batman", "1")
		assert.Nil(t, err)
	})
}
