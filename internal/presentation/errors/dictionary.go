package errors

import (
	"net/http"

	"github.com/andreis3/customers-ms/internal/domain/apperrors"
	"github.com/andreis3/customers-ms/internal/infra/commons/infraerrors"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   int
}

var ErrorDictionary = map[apperrors.ErrorCode]ProtocolError{
	apperrors.ErrInvalidBusinessRules: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   3,
	},
	apperrors.ErrResourceNotFound: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   5,
	},
	infraerrors.ErrInternalProcessing: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   13,
	},
}
