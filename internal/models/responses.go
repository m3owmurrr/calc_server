package models

import "errors"

var (
	ErrNotValidJson       = errors.New("sended data is invalid")
	ErrNotValidExpression = errors.New("expression is not valid")
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ResultResponse struct {
	Result float64 `json:"result"`
}
