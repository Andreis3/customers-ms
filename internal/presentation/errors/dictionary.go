package errors

import (
	"net/http"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   int
}

var ErrorDictionary = map[apperror.Code]ProtocolError{
	apperror.ErrInvalidBusinessRules: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   3,
	},
	apperror.ErrResourceNotFound: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   5,
	},
	apperror.ErrInternalProcessing: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   13,
	},
}
