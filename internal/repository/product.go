package repository

import (
	"github.com/efmatui/product-api/internal/datastruct"
)

func GetAllProduct() ([]datastruct.Product, error) {

	var products []datastruct.Product
	err := NewDB().Model(&products).Select()
	if err != nil {
		panic(err)
	}

	return products, nil
}
