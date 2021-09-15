package main

import (
	"api/internal/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := SetupRouter()
	log.Fatal(router.Run(":8080"))
}

func SetupRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/book/:id", controllers.GetBook)
		v1.POST("/book", controllers.CreateNewBook)
		v1.GET("/books", controllers.GetAllBooks)
		v1.DELETE("/book", controllers.DeleteBook)
		v1.DELETE("/book/:id", controllers.DeleteBook)
		v1.GET("/check", controllers.HealthCheck)
	}
	return router
}
