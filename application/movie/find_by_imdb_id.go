package movie

import (
	"backend-go-recruitment/common/movie"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (ma *movieApplication) FindMovieByImdbID(imdbID string) (*movie.MovieDetail, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", apiKey, imdbID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := ma.client.Do(request)
	if err != nil {
		log.Println("error do request, err: ", err.Error())
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("error read body, err: ", err.Error())
		return nil, err
	}

	finalResp := movie.MovieDetail{}
	err = json.Unmarshal(body, &finalResp)
	if err != nil {
		log.Println("error unmarshal, err: ", err.Error())
		return nil, err
	}

	return &finalResp, err
}
