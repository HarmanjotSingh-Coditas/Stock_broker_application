package routes

import (
	"authentication/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	router.POST("/Signup", handlers.SignUpHandler)
}
