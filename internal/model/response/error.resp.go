package response

import (
	"errors"
	"strings"

	"github.com/bagastri07/gin-boilerplate-service/internal/common/constant"
	"github.com/bagastri07/gin-boilerplate-service/internal/model/customerror"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ErrorResp struct {
	Error *customerror.CustomError `json:"error"`
}

func HandleError(c *gin.Context, err error) {
	var (
		fErr *customerror.CustomError
	)

	switch cErr := err.(type) {
	case *customerror.CustomError:
		fErr = cErr
	case validator.ValidationErrors:
		fErr = processValidationErr(cErr)
	default:
		fErr = processDefaultErr(err)
	}

	c.JSON(fErr.Code, ErrorResp{
		Error: fErr,
	})
}

func processDefaultErr(err error) *customerror.CustomError {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return constant.ErrDataNotFound
	}

	for key, value := range constant.MapErrToCustomError {
		if strings.Contains(err.Error(), key) {
			return value
		}
	}

	return constant.ErrInternalServerError
}

func processValidationErr(ginErr validator.ValidationErrors) *customerror.CustomError {
	var validationErrors []customerror.ValidationError

	for _, fieldError := range ginErr {
		validationError := customerror.ValidationError{
			Field:   fieldError.Field(),
			Tag:     fieldError.Tag(),
			Message: fieldError.Error(),
		}
		validationErrors = append(validationErrors, validationError)
	}

	errorResponse := *constant.ErrBindingValidationError
	errorResponse.ValidationError = validationErrors

	return &errorResponse
}
