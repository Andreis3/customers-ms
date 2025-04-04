package errors

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/errors"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   int
}

var ErrorDictionary = map[errors.ErrorCode]ProtocolError{
	errors.ErrInvalidBusinessRules: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   3,
	},
	errors.ErrResourceNotFound: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   5,
	},
	errors.ErrInternalProcessing: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   13,
	},
}
