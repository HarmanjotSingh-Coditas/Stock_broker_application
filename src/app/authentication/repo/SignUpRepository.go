package repo

import (
	"authentication/constants"
	"errors"
	dbmodels "stock_broker_application/models"
	"stock_broker_application/src/utils/db"

	"gorm.io/gorm"
)

func CheckUserExists(email, username string) (bool, error) {
	var user dbmodels.User
	err := db.DB.Where("email = ? OR username = ?", email, username).First(&user).Error
	if err == nil {
		return true, errors.New(constants.ErrUserExists)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}
func CreateUser(user dbmodels.User) error {
	if err := db.DB.Create(&user).Error; err != nil {
		return errors.New(constants.ErrCreateUser)
	}
	return nil
}
