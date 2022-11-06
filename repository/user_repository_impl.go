package repository

import (
	"errors"

	"github.com/krifik/test_fullstack_backend/config"
	"github.com/krifik/test_fullstack_backend/entity"
	"github.com/krifik/test_fullstack_backend/exception"
	"github.com/krifik/test_fullstack_backend/helper"
	"github.com/krifik/test_fullstack_backend/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repository *UserRepositoryImpl) FindAll() ([]entity.User, error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()

	var users []entity.User
	result := repository.DB.WithContext(ctx).Find(&users)

	if result.RowsAffected < 0 {
		return nil, errors.New("User not found")
	}

	return users, nil
}

func (repository *UserRepositoryImpl) Delete(db *gorm.DB, userId int) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result := db.WithContext(ctx).Delete(&userId)
	exception.PanicIfNeeded(result.Error)
}
func (repository *UserRepositoryImpl) Register(request model.CreateUserRequest) (user entity.User, err error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	request.Password = helper.ToHashedPassword(request.Password)
	user = entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
	result := repository.DB.WithContext(ctx).Create(&user)
	exception.PanicIfNeeded(result.Error)
	return user, err
}

func (repository *UserRepositoryImpl) Login(request model.CreateUserRequest) (user entity.User, err error) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	err = repository.DB.WithContext(ctx).First(&user, "email=?", request.Email).Error
	if err != nil {
		return user, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return user, err
	}
	return user, err
}

func (repository *UserRepositoryImpl) CheckEmail(request model.CreateUserRequest) (result int64) {
	ctx, cancel := config.NewPostgresContext()
	defer cancel()
	result = repository.DB.WithContext(ctx).First(&entity.User{}, "email=?", request.Email).RowsAffected
	return result
}
