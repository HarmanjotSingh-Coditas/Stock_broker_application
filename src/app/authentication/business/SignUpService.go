package business

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repo"
	"errors"
	dbmodels "stock_broker_application/models"
	"stock_broker_application/src/utils"
)

func SignUpService(bffCreateUserRequest models.BFFUserRequest) (string, error) {
	exists, err := repo.CheckUserExists(bffCreateUserRequest.EmailId, bffCreateUserRequest.UserName)
	if err != nil {
		return "", err
	}
	if exists {
		return "", errors.New(constants.ErrUserExists)
	}

	hashedPassword, err := utils.HashPassword(bffCreateUserRequest.Password)
	if err != nil {
		return "", err
	}

	newUser := dbmodels.User{
		UserName:    bffCreateUserRequest.UserName,
		EmailID:     bffCreateUserRequest.EmailId,
		PhoneNumber: bffCreateUserRequest.PhoneNumber,
		PANNumber:   bffCreateUserRequest.PANNumber,
		Password:    hashedPassword,
	}

	err = repo.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return constants.MessageUserCreated, nil
}
