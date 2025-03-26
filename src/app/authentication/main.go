package main

import (
	"authentication/routes"
	"log"
	"stock_broker_application/src/utils/db"
	"stock_broker_application/src/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()
	database := db.InitDB() 
	router := gin.Default()

	routes.InitializeRoutes(router, database)

	log.Println("Starting Server on :8080")
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatalf("Server error: %s", err.Error())
	}
}
