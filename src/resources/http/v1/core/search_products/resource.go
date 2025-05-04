package search_products

import (
	searchProductsAggregator "core/src/services/core/search_products_aggregator"
	"github.com/go-playground/form"
)

type Resource struct {
	service     searchProductsAggregator.Aggregator
	formDecoder *form.Decoder
}

func NewResource(service searchProductsAggregator.Aggregator) Resource {
	return Resource{service: service, formDecoder: form.NewDecoder()}
}

//func (res Resource) Routes() chi.Router {
//	router := chi.NewRouter()
//
//	router.Route("/products", func(r chi.Router) {
//		r.Get("/search", res.SearchProducts)
//	})
//
//	return router
//}
