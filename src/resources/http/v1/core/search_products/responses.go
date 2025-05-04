package search_products

import (
	searchProductsAggregator "core/src/services/core/search_products_aggregator"
	"time"
)

type Response struct {
	Result SearchProductsResults `json:"result"`
}

type SearchProductsResults []SearchProductsResult

type SearchProductsResult struct {
	Source   string   `json:"source"`
	Products Products `json:"products"`
}

type Products []Product

type Product struct {
	ID                 string    `json:"id"`                  // Артикул или другой идентификатор
	Name               string    `json:"name"`                // Наименование продукта
	URL                string    `json:"url"`                 // Ссылка на страницу продукта (pdp)
	ImageURL           string    `json:"image_url"`           // Ссылка на изображение продукта
	CurrentPrice       int       `json:"current_price"`       // Текущая цена, по которой продают
	OriginalPrice      int       `json:"original_price"`      // Базовая цена, от которой считается скидка
	DiscountPercentage int       `json:"discount_percentage"` // Скидка в процентах
	ReviewsCount       int       `json:"reviews_count"`       // Количество отзывов
	Rating             float64   `json:"rating"`              // Рейтинг продукта
	DeliveryDate       time.Time `json:"delivery_date"`       // Дата доставки
}

func GetResponse(searchProductsResponse searchProductsAggregator.SearchProductsResponse) Response {
	searchProductsResults := make(SearchProductsResults, 0, len(searchProductsResponse.SearchProductsResults))
	for _, results := range searchProductsResponse.SearchProductsResults {
		products := make(Products, 0, len(results.Products))
		for _, resultProduct := range results.Products {
			product := Product{
				ID:                 resultProduct.ID,
				Name:               resultProduct.Name,
				URL:                resultProduct.URL,
				ImageURL:           resultProduct.ImageURL,
				CurrentPrice:       resultProduct.CurrentPrice,
				OriginalPrice:      resultProduct.OriginalPrice,
				DiscountPercentage: resultProduct.DiscountPercentage,
				ReviewsCount:       resultProduct.ReviewsCount,
				Rating:             resultProduct.Rating,
				DeliveryDate:       resultProduct.DeliveryDate,
			}
			products = append(products, product)
		}

		searchProductsResult := SearchProductsResult{
			Source:   string(results.Source),
			Products: products,
		}

		searchProductsResults = append(searchProductsResults, searchProductsResult)
	}

	response := Response{
		Result: searchProductsResults,
	}

	return response
}
