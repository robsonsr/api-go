package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(productUseCase usecase.ProductUseCase) productController {
	return productController{
		productUseCase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUseCase.GetProducts()

	fmt.Println(err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.NewInternalServerError(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, products)
}
