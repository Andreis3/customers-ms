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
	RequestID    string `json:"request_id"`
	StatusCode   int    `json:"status_code"`
	CodeError    string `json:"code_error"`
	ErrorMessage any    `json:"error_message"`
}

type TypeResponseSuccess struct {
	RequestID  string `json:"request_id"`
	StatusCode int    `json:"status_code"`
	Data       any    `json:"data"`
}

func ResponseSuccess[T any](write http.ResponseWriter, requestID string, status int, data T) {
	write.Header().Set(ContentType, ApplicationJSON)
	write.WriteHeader(status)
	result := TypeResponseSuccess{
		RequestID:  requestID,
		StatusCode: status,
		Data:       data,
	}
	_ = json.NewEncoder(write).Encode(result)
}

func ResponseError[T any](write http.ResponseWriter, status int, requestID, codeError string, data T) {
	write.Header().Set(ContentType, ApplicationJSON)
	write.WriteHeader(status)
	result := TypeResponseError{
		RequestID:    requestID,
		CodeError:    codeError,
		StatusCode:   status,
		ErrorMessage: data,
	}
	_ = json.NewEncoder(write).Encode(result)
}
