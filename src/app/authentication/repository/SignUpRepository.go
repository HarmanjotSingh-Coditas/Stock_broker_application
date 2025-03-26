package repository

import (
	"authentication/constants"
	"errors"
	dbmodels "stock_broker_application/src/models"

	"gorm.io/gorm"
)

type UserSignUpRepository struct {
	DB *gorm.DB
}

func NewUserSignUpRepository(db *gorm.DB) *UserSignUpRepository {
	return &UserSignUpRepository{DB: db}
}

func (repo *UserSignUpRepository) CheckUserExists(email, username , pan string) (bool, error) {
	var user dbmodels.User
	err := repo.DB.Where("email = ? OR username = ? OR PAN = ?", email, username , pan).First(&user).Error

	if err == nil {
		return true, errors.New(constants.ErrUserExists)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}

func (repo *UserSignUpRepository) CreateUser(user dbmodels.User) error {
	if err := repo.DB.Create(&user).Error; err != nil {
		return errors.New(constants.ErrCreateUser)
	}
	return nil
}
