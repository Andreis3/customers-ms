package errors

import (
	"net/http"

	"github.com/andreis3/users-ms/internal/domain/errors"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   int // (exemplo simplificado)
}

var ErrorDictionary = map[errors.ErrorCode]ProtocolError{
	errors.ErrInvalidBusinessRules: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   3, // INVALID_ARGUMENT (código gRPC real)
	},
	errors.ErrResourceNotFound: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   5, // NOT_FOUND
	},
	errors.ErrInternalProcessing: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   13, // INTERNAL
	},
}
