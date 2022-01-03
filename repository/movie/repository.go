package movie

import "database/sql"

type MovieRepository interface {
	Insert(keyword, pagination string) error
}

type movieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) MovieRepository {
	return &movieRepository{
		db: db,
	}
}
