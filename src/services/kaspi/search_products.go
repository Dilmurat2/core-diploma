package kaspi

import (
	"context"
	kaspiAdapter "core/src/adapters/kaspi"
	kaspiClients "core/src/clients/kaspi"
	"core/src/models"
	"core/src/services"
	"fmt"
	"log"
)

type SearchProductsService struct {
	client                  kaspiClients.Client
	productsResponseAdapter kaspiAdapter.ProductsResponseAdapter
}

func NewSearchProductsService(
	client kaspiClients.Client,
	productsResponseAdapter kaspiAdapter.ProductsResponseAdapter,
) SearchProductsService {
	return SearchProductsService{
		client:                  client,
		productsResponseAdapter: productsResponseAdapter,
	}
}

func (s SearchProductsService) SearchProducts(ctx context.Context, request services.SearchProductsServiceRequest) (models.Products, error) {
	limit := 10

	getProductsRequest := kaspiClients.GetProductsRequest{
		Query: request.Query,
		Sort:  "relevance",
		Limit: limit,
		Page:  1,
	}

	response, err := s.client.GetProducts(ctx, getProductsRequest)
	if err != nil {
		log.Printf(fmt.Sprintf("[ERROR] can not search, err: %s", err))
		return models.Products{}, err
	}

	products := s.productsResponseAdapter.GetProducts(response, limit)

	return products, nil
}
