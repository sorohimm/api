package main

import (
	"api/internal/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//router := SetupRouter()
	router := gin.Default()

	router.GET("/book/:id", controllers.GetBook)
	router.POST("/book", controllers.CreateNewBook)
	router.GET("/books", controllers.GetAllBooks)
	router.DELETE("/book", controllers.DeleteBook)
	router.DELETE("/book/:id", controllers.DeleteBook)
	router.GET("/check", controllers.HealthCheck)

	log.Fatal(router.Run())
}

/*func SetupRouter() *gin.Engine {

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
}*/
