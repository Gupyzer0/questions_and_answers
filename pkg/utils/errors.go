package utils

import "net/http"

type CustomErrorInt interface{
	Error()
	GetCode()
}

type CustomError struct {
	message string
	code    int
}

func (err CustomError) Error() string {
	return err.message
}

func (err CustomError) GetCode() int {
	return err.code
}

var ErrNotFound = CustomError{
	message: "Resource Not Found",
	code: http.StatusNotFound,
}

var ErrBadData = CustomError{
	message: "Bad Request",
	code: http.StatusBadRequest,
}
var ServerError = CustomError{
	message: "Server Error",
	code: http.StatusInternalServerError,
}
