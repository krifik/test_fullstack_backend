package faker

import (
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/helper"

	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *entity.User {
	return &entity.User{
		Email:    "test@62teknologi.com",
		Password: helper.ToHashedPassword("Pass1234"),
	}
}
