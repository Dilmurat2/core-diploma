package wb

type GetProductsRequest struct {
	Query       string
	Limit       int
	Page        int
	Currency    string // Валюта
	Language    string // Язык
	Destination string // Значение меняется в зависимости от выбранной точки самовывоза
	Sort        string
}
