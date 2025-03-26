package routes

import (
	"authentication/business"
	"authentication/handlers"
	"authentication/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
	signInRepo := repository.NewUserSignInRepository(db)
	signInService := business.NewSignInService(signInRepo)
	signInController := handlers.NewSignInController(signInService)

	signUpRepo := repository.NewUserSignUpRepository(db)
	signUpService := business.NewSignUpService(signUpRepo)
	signUpController := handlers.NewSignUpController(signUpService)

	router.POST("/Signup", signUpController.SignUpHandler)
	router.POST("/Signin", signInController.SignInHandler)
}
