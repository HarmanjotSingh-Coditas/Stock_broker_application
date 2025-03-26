package business

import (
	"authentication/constants"
	"authentication/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type SignInService struct {
	repository *repository.UserSignInRepository
}

func NewSignInService(repo *repository.UserSignInRepository) *SignInService {
	return &SignInService{repository: repo}
}

func (service *SignInService) SignIn(email, password string) error {
	user, err := service.repository.GetUserByEmail(email)
	if err != nil {
		return errors.New(constants.ErrInvalidCredentials)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return errors.New(constants.ErrInvalidCredentials)
	}

	return nil
}
