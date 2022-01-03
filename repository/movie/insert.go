package movie

import (
	"log"
	"time"
)

var query = `INSERT INTO SEARCH_LOGS (keyword, page, timestamp) VALUES (?, ?, ?)`

func (mr *movieRepository) Insert(keyword, pagination string) error {
	statement, err := mr.db.Prepare(query)
	if err != nil {
		log.Println("error preparing statement, error:", err.Error())
		return err
	}

	_, err = statement.Exec(keyword, pagination, time.Now())
	if err != nil {
		log.Println("error exec statement, error:", err.Error())
	}
	return err
}
