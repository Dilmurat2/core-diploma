package search_products

import (
	"core/src/resources/http/schema"
	searchProductsAggregator "core/src/services/core/search_products_aggregator"
	"github.com/go-chi/render"
	"net/http"
)

func (res Resource) SearchProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var queryParams SearchProductsRequest

	if err := res.formDecoder.Decode(&queryParams, r.URL.Query()); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{"error": "invalid query parameters"})
		return
	}

	searchProductsRequest := searchProductsAggregator.SearchProductsAggregatorRequest{
		Query: queryParams.Query,
	}

	aggregateProducts, err := res.service.SearchProducts(ctx, searchProductsRequest)
	if err != nil {
		_ = render.Render(w, r, schema.BadRequest(err))
		return
	}

	response := GetResponse(aggregateProducts)

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
