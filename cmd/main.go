package main

import (
	"api/internal/config"
	"api/internal/controllers"
	"api/internal/infrastructure"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"os"
)

var (
	cfg *config.Config
	ctx context.Context
	log *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Printf("error loading logger: %s", err)
		os.Exit(1)
		return
	}

	log = logger.Sugar()

	cfg, err = config.New()
	if err != nil {
		log.Fatalf("config init error: %s", err)
	}
	log.Infof("Config loaded:\n%+v", cfg)

	ctx = context.Background()
}

func main() {
	injector, _ := infrastructure.Injector(log, ctx, cfg)
	bookController := injector.InjectBookController()

	router := gin.Default()

	v1 := router.Group("/books/v1")
	{
		v1.GET("/book/:id", bookController.GetBook)
		v1.POST("/book", bookController.CreateBook)
		// TODO: put request
		v1.GET("/books", bookController.GetAllBooks)
		v1.DELETE("/book/:id", controllers.DeleteBook)
		v1.GET("/check", controllers.HealthCheck)
	}

	log.Fatal(router.Run())
}
