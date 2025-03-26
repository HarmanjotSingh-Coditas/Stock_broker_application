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
	signUpRepo := repository.NewUserSignUpRepository(db)

	signInService := business.NewSignInService(signInRepo)
	signUpService := business.NewSignUpService(signUpRepo)

	signInController := handlers.NewSignInController(signInService)
	signUpController := handlers.NewSignUpController(signUpService)

	router.POST("/Signup", signUpController.SignUpHandler)
	router.POST("/Signin", signInController.SignInHandler)
}
