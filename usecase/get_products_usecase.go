package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type GetProductsUseCase struct {
	repository repository.ProductRepository
}

func NewGetProductsUseCase(repository repository.ProductRepository) GetProductsUseCase {
	return GetProductsUseCase{
		repository,
	}
}

func (p *GetProductsUseCase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}
