package datastruct

const ProductTableName = "products"

type Product struct {
	Id          int64
	Name        string
	Price       int64
	Description string
	Quantity    int64
	Visible     bool
}
