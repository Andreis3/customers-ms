package errors

import (
	"net/http"

	domain_errors "github.com/andreis3/users-ms/internal/domain/errors"
	infra_errors "github.com/andreis3/users-ms/internal/infra/commons/errors"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   int
}

var ErrorDictionary = map[domain_errors.ErrorCode]ProtocolError{
	domain_errors.ErrInvalidBusinessRules: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   3,
	},
	domain_errors.ErrResourceNotFound: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   5,
	},
	infra_errors.ErrInternalProcessing: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   13,
	},
}
