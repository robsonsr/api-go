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
	GetProductsUseCase := usecase.NewGetProductsUseCase(ProductRepository)
	CreateProductUseCase := usecase.NewCreateProductsUseCase(ProductRepository)

	//Camada de controllers
	ProductController := controller.NewProductController(GetProductsUseCase, CreateProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]string{
			"message": "pongggg",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8000")
}
