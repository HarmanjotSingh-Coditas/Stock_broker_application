package handlers

import (
	"authentication/business"
	"authentication/constants"
	"authentication/models"
	"net/http"
	"reflect"
	"stock_broker_application/src/utils/validators"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignUpHandler(c *gin.Context) {
	var bffCreateUserRequest models.BFFUserRequest

	if err := c.ShouldBindJSON(&bffCreateUserRequest); err != nil {
		var validationErrors []string

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				field, _ := reflect.TypeOf(bffCreateUserRequest).FieldByName(fieldErr.StructField())
				jsonTag := field.Tag.Get("json")

				validationErrors = append(validationErrors, jsonTag+" is required")
			}
		} else {
			validationErrors = append(validationErrors, constants.ErrInvalidRequestBody)
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}
	if bffCreateUserRequest.Password != bffCreateUserRequest.ConfirmPassword {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: constants.ErrPasswordMismatch,
		})
		return
	}

	if !validators.ValidatePAN(bffCreateUserRequest.PANNumber) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: constants.ErrInvalidPAN,
		})
		return
	}
	if !validators.ValidatePhone(bffCreateUserRequest.PhoneNumber) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: constants.ErrInvalidPhone,
		})
		return
	}
	if !validators.ValidateEmail(bffCreateUserRequest.EmailId) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: constants.ErrInvalidEmail,
		})
		return
	}
	if !validators.ValidatePassword(bffCreateUserRequest.Password) {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: constants.ErrInvalidPassword,
		})
		return
	}

	message, err := business.SignUpService(bffCreateUserRequest)
	if err != nil {
		if err.Error() == constants.ErrUserExists {
			c.JSON(http.StatusConflict, models.ErrorResponse{
				Error: constants.ErrUserExists,
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.BFFUserResponse{
		Message: message,
	})
}
