package main

import (
	"fmt"
	"log"

	"authenticator/controllers"
	"authenticator/models"
	"authenticator/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB              *gorm.DB
	server          *gin.Engine
	Controller      controllers.Controller
	RouteController routes.RouteController
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("store.db"), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("failed to connect to the database")
	}
	fmt.Println("connected successfully to the database")

	Controller = controllers.NewController(DB)
	RouteController = routes.NewRouteController(Controller)
	server = gin.Default()
}

func main() {
	router := server.Group("/api")
	RouteController.AuthRoute(router)
	log.Fatal(server.Run(":5000"))
}
