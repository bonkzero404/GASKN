package http

import "fmt"

type SetResponse struct {
	Valid bool    `json:"valid"`
	Meta  SetMeta `json:"meta"`
	Error any     `json:"errors"`
	Data  any     `json:"data"`
}

type SetMeta struct {
	Route  string `json:"route"`
	Method string `json:"method"`
	Query  string `json:"query"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type SetErrors struct {
	Message string `json:"message"`
	Cause   string `json:"cause"`
	Inputs  any    `json:"inputs"`
}

type SetErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Message     string `json:"message"`
}

type SetApiErrorResponse struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (r *SetApiErrorResponse) Error() string {
	// return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Message)
	return fmt.Sprintf("%v", r.Message)
}

func (r *SetApiErrorResponse) GetStatusCode() int {
	return r.StatusCode
}
