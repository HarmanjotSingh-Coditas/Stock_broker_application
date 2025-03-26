package repository

import (
	"authentication/constants"
	"errors"
	dbmodels "stock_broker_application/src/models"

	"gorm.io/gorm"
)

type UserSignInRepository struct {
	DB *gorm.DB
}

func NewUserSignInRepository(db *gorm.DB) *UserSignInRepository {
	return &UserSignInRepository{DB: db}
}

func (repo *UserSignInRepository) GetUserByEmail(email string) (*dbmodels.User, error) {
	var user dbmodels.User
	result := repo.DB.Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(constants.ErrUserNotFound)
	} else if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
