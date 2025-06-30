package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/error"
)

func DecoderBodyRequest[T any](req *http.Request) (T, *error.Error) {
	defer req.Body.Close()
	var result T
	var jsonUnmarshalTypeError *json.UnmarshalTypeError
	var jsonSyntaxError *json.SyntaxError
	err := json.NewDecoder(req.Body).Decode(&result)
	switch {
	case errors.As(err, &jsonSyntaxError):
		return result, ErrorJSONSyntaxError(jsonSyntaxError)

	case errors.As(err, &jsonUnmarshalTypeError):
		return result, ErrorJSONUnmarshalTypeError(jsonUnmarshalTypeError)

	case err != nil:
		return result, ErrorJSON(err)
	}

	return result, nil
}

func ErrorJSONSyntaxError(err error) *error.Error {
	return &error.Error{
		Code:            error.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json syntax error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json syntax error",
	}
}

func ErrorJSONUnmarshalTypeError(err error) *error.Error {
	return &error.Error{
		Code:            error.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json unmarshal type error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json unmarshal type error",
	}
}

func ErrorJSON(err error) *error.Error {
	return &error.Error{
		Code:            error.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json error",
	}
}
