package handlers

import (
	"authentication/business"
	"authentication/constants"
	"authentication/models"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

type SignInController struct {
	service *business.SignInService
}

func NewSignInController(service *business.SignInService) *SignInController {
	return &SignInController{service: service}
}

func (controller *SignInController) SignInHandler(c *gin.Context) {
	var bffUserSignInRequest models.BFFUserSignInRequest

	if err := c.ShouldBindJSON(&bffUserSignInRequest); err != nil {
		var validationErrors []map[string]string

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				field, _ := reflect.TypeOf(bffUserSignInRequest).FieldByName(fieldErr.StructField())
				jsonTag := field.Tag.Get("json")

				validationErrors = append(validationErrors, map[string]string{
					"key":   jsonTag,
					"error": jsonTag + " is required",
				})
			}
		} else {
			validationErrors = append(validationErrors, map[string]string{
				"key":   "request",
				"error": constants.ErrInvalidRequestBody,
			})
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	err := controller.service.SignIn(bffUserSignInRequest.EmailId, bffUserSignInRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.BFFUserResponse{
		Message: constants.LoginSuccessfull,
	})
}
