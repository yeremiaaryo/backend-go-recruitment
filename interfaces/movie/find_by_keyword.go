package movie

import (
	"backend-go-recruitment/common/movie"
	"errors"
	"log"
	"sync"
)

func (mi *movieInterface) FindMovieByKeyword(keyword, pagination string) (*movie.MovieList, error) {
	resp, err := mi.movieApplication.FindMovieByKeyword(keyword, pagination)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, errors.New("empty response from OMDB API")
	}

	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)

	wg.Add(len(resp.Search))

	movieMap := make(map[string]movie.Movie)
	for _, v := range resp.Search {
		go func(m movie.Movie) {
			detail, err := mi.movieApplication.FindMovieByImdbID(m.ImdbID)
			if err != nil {
				log.Println("error when find movie by imdbID with imdbID:", m.ImdbID)
			}
			mu.Lock()
			m.Detail = detail
			movieMap[m.ImdbID] = m
			mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()

	for i, v := range resp.Search {
		mv := movieMap[v.ImdbID]
		resp.Search[i] = mv
	}

	err = mi.movieApplication.InsertLogs(keyword, pagination)
	if err != nil {
		log.Println("error inserting logs to database", err.Error())
	}
	return resp, nil
}
