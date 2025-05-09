package main

import (
	"swag-gin-demo/handlers"

	"github.com/gin-gonic/gin"

	_ "swag-gin-demo/docs"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/kwinwithaki/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @BasePath /
// @query.collection.format multi

func main() {
	// configure the Gin server
	router := gin.Default()

	router.GET("/todo", handlers.GetAllTodos)
	router.GET("/todo/:id", handlers.GetTodoByID)

	// docs routes
	router.StaticFile("/swagger.json", "./docs/swagger.json")
	router.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
	router.StaticFile("/", "index.html")

	// run the Gin server
	router.Run()
}
