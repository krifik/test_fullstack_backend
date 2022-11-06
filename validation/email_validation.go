package validation

import "github.com/krifik/test_fullstack_backend/exception"

func IsEmailHasBeenTaken(row int64) {
	if row != 0 {
		panic(exception.UniqueEmailError{
			Message: "Email has been taken",
		})
	}
}
