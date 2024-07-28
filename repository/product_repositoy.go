package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	dbConnection *sql.DB
}

func NewProductRepository(dbConnection *sql.DB) ProductRepository {
	return ProductRepository{
		dbConnection,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"

	rows, err := p.dbConnection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObject model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObject.Id,
			&productObject.Name,
			&productObject.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObject)

	}

	rows.Close()

	return productList, nil
}

func (p *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int

	query, err := p.dbConnection.Prepare("INSERT INTO product(product_name, price) VALUES($1, $2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, nil
}
