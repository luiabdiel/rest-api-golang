package main

import (
	"rest-api-golang/controller"
	"rest-api-golang/db"
	"rest-api-golang/repository"
	"rest-api-golang/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Camada de usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	// Camada de contrllers
	productController := controller.NewProductController(ProductUsecase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", productController.GetProducts)

	server.Run(":8000")
}
