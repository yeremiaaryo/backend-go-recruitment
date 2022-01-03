package common

import "net/http"

type BaseResponseDto struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
