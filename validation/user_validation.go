package validation

import (
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func Validate(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		// validation.Field(&request.Id, validation.Required),
		// validation.Field(&request.FirstName, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required, validation.Length(5, 10)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
