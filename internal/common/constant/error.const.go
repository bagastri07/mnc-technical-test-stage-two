package constant

import (
	"net/http"

	"github.com/bagastri07/mnc-technical-test-stage-two/internal/model/customerror"
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
	ErrUnauthenticated = &customerror.CustomError{
		Code:    http.StatusUnauthorized,
		Message: "Unauthenticated",
	}
	ErrPhoneNumberAlreadyRegistered = &customerror.CustomError{
		Code:    http.StatusBadRequest,
		Message: "Phone Number already registered",
	}
	ErrPhoneNumberAndPINDoesntMatch = &customerror.CustomError{
		Code:    http.StatusUnauthorized,
		Message: "Phone Number and PIN doesn’t match",
	}
	ErrBalanceIsNotEnough = &customerror.CustomError{
		Code:    http.StatusBadRequest,
		Message: "Balance is not enough",
	}
)

var (
	MapErrToCustomError = map[string]*customerror.CustomError{}
)
