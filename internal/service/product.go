package service

import (
	"log"

	"github.com/efmatui/product-api/internal/datastruct"
	"github.com/efmatui/product-api/internal/repository"
)

type ProductService interface {
	GetAllProduct() ([]datastruct.Product, error)
}

type productService struct {
	dao repository.DAO
}

func (u *productService) GetAllProduct() ([]datastruct.Product, error) {
	var products []datastruct.Product
	var err error
	products, err = u.dao.NewProductQuery().GetAllProduct()
	if err != nil {
		log.Printf("error products %v", err)
	}

	return products, err
}
