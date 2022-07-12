package service

type DAO interface {
	NewProductService() ProductService
}

type dao struct{}

func (d *dao) NewProductService() ProductService {
	return &productService{}
}
