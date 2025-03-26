package business

import (
	"authentication/constants"
	"authentication/models"
	"authentication/repository"
	"errors"
	dbmodels "stock_broker_application/src/models"
	"stock_broker_application/src/utils"
)

type SignUpService struct {
	userRepo *repository.UserSignUpRepository
}

func NewSignUpService(userRepo *repository.UserSignUpRepository) *SignUpService {
	return &SignUpService{userRepo: userRepo}
}

func (service *SignUpService) SignUpUser(bffCreateUserRequest models.BFFUserRequest) (string, error) {
	exists, err := service.userRepo.CheckUserExists(bffCreateUserRequest.EmailId, bffCreateUserRequest.UserName, bffCreateUserRequest.PANNumber)
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

	err = service.userRepo.CreateUser(newUser)
	if err != nil {
		return "", err
	}
	return constants.MessageUserCreated, nil
}
