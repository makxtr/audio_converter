package apperrors

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Code    int
	Message string
}

func (e HttpError) Error() string {
	return e.Message
}

func NewHttpError(code int, message string) error {
	return HttpError{
		Code:    code,
		Message: message,
	}
}

func WriteHttpError(w http.ResponseWriter, err error) {
	var httpErr HttpError
	if errors.As(err, &httpErr) {
		http.Error(w, httpErr.Message, httpErr.Code)
	} else {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
