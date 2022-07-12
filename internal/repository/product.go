package repository

import (
	"github.com/efmatui/product-api/internal/datastruct"
)

type ProductQuery interface {
	GetAllProduct() ([]datastruct.Product, error)
}

type productQuery struct{}

func (p *productQuery) GetAllProduct() ([]datastruct.Product, error) {

	var products []datastruct.Product
	err := NewDB().Model(&products).Select()
	if err != nil {
		panic(err)
	}

	return products, nil
}
