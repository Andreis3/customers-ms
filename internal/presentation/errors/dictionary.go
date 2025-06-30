package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"

	"github.com/andreis3/customers-ms/internal/domain/error"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   codes.Code
}

var ErrorDictionary = map[error.Code]ProtocolError{
	error.BadRequestCode: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   codes.InvalidArgument,
	},
	error.NotFoundCode: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   codes.NotFound,
	},
	error.InternalServerErrorCode: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   codes.Internal,
	},
	error.UnauthorizedCode: {
		HTTPStatus: http.StatusUnauthorized,
		GRPCCode:   codes.Unauthenticated,
	},
	error.ForbiddenCode: {
		HTTPStatus: http.StatusForbidden,
		GRPCCode:   codes.PermissionDenied,
	},
	error.ConflictCode: {
		HTTPStatus: http.StatusConflict,
		GRPCCode:   codes.AlreadyExists,
	},
	error.UnprocessableEntityCode: {
		HTTPStatus: http.StatusUnprocessableEntity,
		GRPCCode:   codes.InvalidArgument,
	},
}
