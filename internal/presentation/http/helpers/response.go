package helpers

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type TypeResponseError struct {
	CodeError    string `json:"code_error"`
	ErrorMessage any    `json:"error_message"`
}

type TypeResponseSuccess struct {
	Data any `json:"data"`
}

func ResponseSuccess[T any](write http.ResponseWriter, status int, data T) {
	write.Header().Set(ContentType, ApplicationJSON)
	write.WriteHeader(status)
	result := TypeResponseSuccess{
		Data: data,
	}
	_ = json.NewEncoder(write).Encode(result)
}

func ResponseError[T any](write http.ResponseWriter, status int, codeError string, data T) {
	write.Header().Set(ContentType, ApplicationJSON)
	write.WriteHeader(status)
	result := TypeResponseError{
		CodeError:    codeError,
		ErrorMessage: data,
	}
	_ = json.NewEncoder(write).Encode(result)
}
