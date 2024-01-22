package validation

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	Field string
	Msg   string
}

func getMessage(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return "Invalid length (min)"

	}
	return ""
}

func GetErrors(err error) []ApiError {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ApiError, len(ve))
		for i, fe := range ve {
			out[i] = ApiError{fe.Field(), getMessage(fe.Tag())}
		}

		return out
	}

	return make([]ApiError, 0)
}
