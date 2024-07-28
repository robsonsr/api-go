package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	getProductsUseCase   usecase.GetProductsUseCase
	createProductUseCase usecase.CreateProductUseCase
}

func NewProductController(getProductsUseCase usecase.GetProductsUseCase, createProductUseCase usecase.CreateProductUseCase) productController {
	return productController{
		getProductsUseCase,
		createProductUseCase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.getProductsUseCase.GetProducts()

	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.NewBadRequestError(err.Error()))
		return
	}

	createdProduct, err := p.createProductUseCase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, createdProduct)

}
