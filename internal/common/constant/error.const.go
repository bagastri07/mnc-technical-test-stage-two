package constant

import (
	"net/http"

	"github.com/bagastri07/gin-boilerplate-service/internal/model/customerror"
)

var (
	ErrBindingValidationError = &customerror.CustomError{
		Code:    http.StatusBadRequest,
		Message: "Binding Error Validation",
	}
	ErrInternalServerError = &customerror.CustomError{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
	ErrDataNotFound = &customerror.CustomError{
		Code:    http.StatusNotFound,
		Message: "Data Not Found",
	}
	ErrUnauthorized = &customerror.CustomError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	ErrUserAlreadyExist = &customerror.CustomError{
		Code:    http.StatusBadRequest,
		Message: "User Already Exist",
	}
)

var (
	MapErrToCustomError = map[string]*customerror.CustomError{}
)
