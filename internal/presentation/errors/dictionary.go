package errors

import (
	"net/http"

	"google.golang.org/grpc/codes"

	apperror "github.com/andreis3/customers-ms/internal/domain/app-error"
)

type ProtocolError struct {
	HTTPStatus int
	GRPCCode   codes.Code
}

var ErrorDictionary = map[apperror.Code]ProtocolError{
	apperror.BadRequestCode: {
		HTTPStatus: http.StatusBadRequest,
		GRPCCode:   codes.InvalidArgument,
	},
	apperror.NotFoundCode: {
		HTTPStatus: http.StatusNotFound,
		GRPCCode:   codes.NotFound,
	},
	apperror.InternalServerErrorCode: {
		HTTPStatus: http.StatusInternalServerError,
		GRPCCode:   codes.Internal,
	},
	apperror.UnauthorizedCode: {
		HTTPStatus: http.StatusUnauthorized,
		GRPCCode:   codes.Unauthenticated,
	},
	apperror.ForbiddenCode: {
		HTTPStatus: http.StatusForbidden,
		GRPCCode:   codes.PermissionDenied,
	},
	apperror.ConflictCode: {
		HTTPStatus: http.StatusConflict,
		GRPCCode:   codes.AlreadyExists,
	},
	apperror.UnprocessableEntityCode: {
		HTTPStatus: http.StatusUnprocessableEntity,
		GRPCCode:   codes.InvalidArgument,
	},
}
