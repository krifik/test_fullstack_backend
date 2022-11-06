package validation

import (
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/model"

	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateStudent(request model.StudentRequest) {
	err := validation.ValidateStruct(&request,
		// validation.Field(&request.Id, validation.Required),
		// validation.Field(&request.FirstName, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Gender, validation.Required),
		validation.Field(&request.Dob, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
