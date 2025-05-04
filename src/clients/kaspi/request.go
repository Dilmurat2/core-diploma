package kaspi

type GetProductsRequest struct {
	Query string
	Limit int
	Page  int
	Sort  string
}
