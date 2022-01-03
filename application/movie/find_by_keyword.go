package movie

import (
	"backend-go-recruitment/common/movie"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const apiKey = "faf7e5bb&s"

func (ma *movieApplication) FindMovieByKeyword(keyword, pagination string) (*movie.MovieList, error) {
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", apiKey, keyword, pagination)
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

	finalResp := movie.MovieList{}
	err = json.Unmarshal(body, &finalResp)
	if err != nil {
		log.Println("error unmarshal, err: ", err.Error())
		return nil, err
	}

	page, err := strconv.Atoi(pagination)
	if err != nil {
		log.Println("error converting page, err: ", err.Error())
		return nil, err
	}
	totalResult, err := strconv.Atoi(finalResp.TotalResults)
	if err != nil {
		log.Println("error converting total result, err: ", err.Error())
		return nil, err
	}

	finalResp.Page = page
	finalResp.TotalPage = totalResult / len(finalResp.Search)
	if totalResult%len(finalResp.Search) != 0 {
		finalResp.TotalPage++
	}
	return &finalResp, err
}
