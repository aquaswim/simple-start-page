package pkg

import (
	"github.com/go-playground/validator/v10"
	"simple-start-page/internal/error"
)

var validate = validator.New()

type ErrorDetail struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

func ValidateStruct(data interface{}) *error.AppError {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}
	appErr := new(error.AppError)
	appErr.Code = error.ErrCodeValidationFailed
	appErr.Message = "Validation error"
	// show error validation field
	var errDetail []*ErrorDetail
	for _, err := range err.(validator.ValidationErrors) {
		var element ErrorDetail
		element.FailedField = err.StructNamespace()
		element.Tag = err.Tag()
		element.Value = err.Param()
		errDetail = append(errDetail, &element)
	}
	appErr.ErrorDetail = errDetail
	return appErr
}
