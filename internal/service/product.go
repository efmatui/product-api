package service

import (
	"log"

	"github.com/efmatui/product-api/internal/datastruct"
	"github.com/efmatui/product-api/internal/repository"
)

func GetAllProduct() ([]datastruct.Product, error) {
	var products []datastruct.Product
	var err error
	products, err = repository.GetAllProduct()
	if err != nil {
		log.Printf("error products %v", err)
	}

	return products, err
}
