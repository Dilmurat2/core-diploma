package search_products_aggregator

import (
	"context"
	"core/src/services"
	"log"
	"sync"
)

type Aggregator interface {
	SearchProducts(context.Context, SearchProductsAggregatorRequest) (SearchProductsResponse, error)
}

func NewAggregator(searchProductsServices map[SourceType]services.SearchProductsService) Aggregator {
	return aggregator{searchProductsServices: searchProductsServices}
}

type aggregator struct {
	searchProductsServices map[SourceType]services.SearchProductsService
}

func (a aggregator) SearchProducts(ctx context.Context, request SearchProductsAggregatorRequest) (SearchProductsResponse, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(a.searchProductsServices))

	serviceResults := make([]SearchProductsResult, 0, len(a.searchProductsServices))

	serviceRequest := services.SearchProductsServiceRequest{
		Query: request.Query,
	}

	for source, service := range a.searchProductsServices {
		go func() {
			defer wg.Done()

			products, err := service.SearchProducts(ctx, serviceRequest)
			if err != nil {
				log.Printf("[ERROR] can not search products, source: %s, err: %v", source, err)
			}

			result := SearchProductsResult{
				Source:   source,
				Products: products,
			}

			serviceResults = append(serviceResults, result)
		}()
	}

	wg.Wait()

	response := SearchProductsResponse{
		SearchProductsResults: serviceResults,
	}

	return response, nil
}
