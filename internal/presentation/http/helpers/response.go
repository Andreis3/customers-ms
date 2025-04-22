package helpers

import (
	"encoding/json"
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/apperrors"
	"github.com/andreis3/users-ms/internal/presentation/errors"
)

const (
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
)

type TypeResponseError struct {
	FriendlyMessage any `json:"friendlyMessage"`
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

func ResponseError[T any](write http.ResponseWriter, err *apperrors.AppErrors) {
	status := errors.ErrorDictionary[err.Code].HTTPStatus
	write.Header().Set(ContentType, ApplicationJSON)
	write.WriteHeader(status)

	result := TypeResponseError{
		FriendlyMessage: err.FriendlyMessage,
	}

	_ = json.NewEncoder(write).Encode(result)
}
