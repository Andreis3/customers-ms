package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/errors"
)

func DecoderBodyRequest[T any](req *http.Request) (T, *errors.Error) {
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

func ErrorJSONSyntaxError(err error) *errors.Error {
	return &errors.Error{
		Code:            errors.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json syntax error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json syntax error",
	}
}

func ErrorJSONUnmarshalTypeError(err error) *errors.Error {
	return &errors.Error{
		Code:            errors.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json unmarshal type error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json unmarshal type error",
	}
}

func ErrorJSON(err error) *errors.Error {
	return &errors.Error{
		Code:            errors.BadRequestCode,
		Errors:          []string{err.Error()},
		Cause:           "json error",
		OriginFunc:      "json.Unmarshal",
		FriendlyMessage: "json error",
	}
}
