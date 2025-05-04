package ozon

import (
	"context"
	ozonAdapter "core/src/adapters/ozon"
	ozonClient "core/src/clients/ozon"

	"core/src/models"
	"core/src/services"
	"fmt"
	"log"
)

type SearchProductsService struct {
	client                  ozonClient.Client
	productsResponseAdapter ozonAdapter.ProductsResponseAdapter
}

func NewSearchProductsService(
	client ozonClient.Client,
	productsResponseAdapter ozonAdapter.ProductsResponseAdapter,
) SearchProductsService {
	return SearchProductsService{
		client:                  client,
		productsResponseAdapter: productsResponseAdapter,
	}
}

func (s SearchProductsService) SearchProducts(ctx context.Context, request services.SearchProductsServiceRequest) (models.Products, error) {
	limit := 10

	getProductsRequest := ozonClient.GetProductsRequest{
		Query: request.Query,
		Sort:  "score",
		Limit: limit,
	}

	response, err := s.client.GetProducts(ctx, getProductsRequest)
	if err != nil {
		log.Printf(fmt.Sprintf("[ERROR] can not search, err: %s", err))
		return models.Products{}, err
	}

	products := s.productsResponseAdapter.GetProducts(response, limit)

	return products, nil
}
