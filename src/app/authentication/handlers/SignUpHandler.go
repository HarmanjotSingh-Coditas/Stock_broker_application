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

type SignUpController struct {
	service *business.SignUpService
}

func NewSignUpController(service *business.SignUpService) *SignUpController {
	return &SignUpController{service: service}
}
func (controller *SignUpController) SignUpHandler(c *gin.Context) {
	var bffCreateUserRequest models.BFFUserRequest

	if err := c.ShouldBindJSON(&bffCreateUserRequest); err != nil {
		var validationErrors []map[string]string

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				field, _ := reflect.TypeOf(bffCreateUserRequest).FieldByName(fieldErr.StructField())
				jsonTag := field.Tag.Get("json")

				errorMsg, exists := constants.ErrFieldRequired[jsonTag]
				if !exists {
					errorMsg = "This field is required"
				}

				validationErrors = append(validationErrors, map[string]string{
					"key":   jsonTag,
					"error": errorMsg,
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

	// Custom validation errors
	var validationErrors []map[string]string

	if bffCreateUserRequest.Password != bffCreateUserRequest.ConfirmPassword {
		validationErrors = append(validationErrors, map[string]string{
			"key":   "confirmpassword",
			"error": constants.ErrPasswordMismatch,
		})
	}

	if !validators.ValidatePAN(bffCreateUserRequest.PANNumber) {
		validationErrors = append(validationErrors, map[string]string{
			"key":   "pan",
			"error": constants.ErrInvalidPAN,
		})
	}

	if !validators.ValidateEmail(bffCreateUserRequest.EmailId) {
		validationErrors = append(validationErrors, map[string]string{
			"key":   "email",
			"error": constants.ErrInvalidEmail,
		})
	}

	if !validators.ValidatePassword(bffCreateUserRequest.Password) {
		validationErrors = append(validationErrors, map[string]string{
			"key":   "password",
			"error": constants.ErrInvalidPassword,
		})
	}

	if !validators.ValidatePhoneNumber(bffCreateUserRequest.PhoneNumber) {
		validationErrors = append(validationErrors, map[string]string{
			"key":   "phoneno",
			"error": constants.ErrInvalidPhone,
		})
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	// Call the SignUpService method
	message, err := controller.service.SignUpUser(bffCreateUserRequest)
	if err != nil {
		if err.Error() == constants.ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{"errors": []map[string]string{
				{"key": "email", "error": constants.ErrUserExists},
			}})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.BFFUserResponse{Message: message})
}
