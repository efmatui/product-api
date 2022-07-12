package app

type DAO interface {
	NewGetProduct() GetProduct
}

type dao struct{}

func (d *dao) NewGetProduct() GetProduct {
	return &getProduct{}
}
