package util

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func MinIf(fl validator.FieldLevel) bool {
	tagParams := strings.Split(fl.Param(), " ")

	if len(tagParams) != 3 {
		return false
	}

	targetField := fl.Parent().FieldByName(tagParams[0])
	if !targetField.IsValid() {
		return false
	}

	targetValue := tagParams[1]

	minLen, err := strconv.Atoi(tagParams[2])
	if err != nil {
		return false
	}

	if targetField.String() == targetValue {
		switch k := fl.Field().Kind(); k {
		case reflect.Int64:
			return fl.Field().Int() >= int64(minLen)
		default:
			return fl.Field().Len() >= minLen
		}
	}

	return true
}

func MaxIf(fl validator.FieldLevel) bool {
	tagParams := strings.Split(fl.Param(), " ")

	if len(tagParams) != 3 {
		return false
	}

	targetField := fl.Parent().FieldByName(tagParams[0])
	if !targetField.IsValid() {
		return false
	}

	targetValue := tagParams[1]

	maxLen, err := strconv.Atoi(tagParams[2])
	if err != nil {
		return false
	}

	if targetField.String() == targetValue {
		switch k := fl.Field().Kind(); k {
		case reflect.Int64:
			return fl.Field().Int() <= int64(maxLen)
		default:
			return fl.Field().Len() <= maxLen
		}
	}

	return true
}

// AddValidation Register validator for struct
func AddValidation() {
	// register the validation in main function
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("Min_if", MinIf)
		_ = v.RegisterValidation("Max_if", MaxIf)
	}
}
