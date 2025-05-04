package search_products_aggregator

import "core/src/models"

type SearchProductsResponse struct {
	SearchProductsResults []SearchProductsResult
}

type SearchProductsResult struct {
	Source   SourceType
	Products models.Products
}
