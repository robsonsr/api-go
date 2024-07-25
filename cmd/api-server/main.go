package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}

	//Camada de Repository
	ProductRepository := repository.NewProductRepository(dbConnection)
	//Camada de UseCase
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	//Camada de controllers
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"message": "pongggg",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
