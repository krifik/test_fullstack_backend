package validation

import (
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func AuthValidate(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
