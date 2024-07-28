package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type CreateProductUseCase struct {
	repository repository.ProductRepository
}

func NewCreateProductsUseCase(repository repository.ProductRepository) CreateProductUseCase {
	return CreateProductUseCase{
		repository,
	}
}

func (p *CreateProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	id, err := p.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.Id = id

	return product, err
}
