package movie

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func Test_movieRepository_Insert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("error prepare query", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		query := `INSERT INTO SEARCH_LOGS (keyword, page, timestamp) VALUES (?, ?, ?)`
		mock.ExpectPrepare(query).WillReturnError(errors.New("failed"))

		mr := movieRepository{db: db}
		err = mr.Insert("batman", "1")
		assert.Error(t, err)
	})

	t.Run("error exec query", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		query := `INSERT INTO SEARCH_LOGS (keyword, page, timestamp) VALUES (?, ?, ?)`
		mock.ExpectPrepare(query).ExpectExec().WithArgs("batman", "1", AnyTime{}).WillReturnError(errors.New("failed"))

		mr := movieRepository{db: db}
		err = mr.Insert("batman", "1")
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		query := `INSERT INTO SEARCH_LOGS (keyword, page, timestamp) VALUES (?, ?, ?)`
		mock.ExpectPrepare(query).ExpectExec().WithArgs("batman", "1", AnyTime{}).WillReturnResult(sqlmock.NewResult(1, 1))

		mr := movieRepository{db: db}
		err = mr.Insert("batman", "1")
		assert.Nil(t, err)
	})
}
