package main

import (
	"authentication/constants"
	"authentication/routes"
	"log"
	"stock_broker_application/src/utils"
	"stock_broker_application/src/utils/db"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()
	db.InitDB()
	router := gin.Default()
	routes.InitializeRoutes(router)
	log.Println("Starting Server on :8080")
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Server error: %s", constants.ErrInternalServer)
	}
}
