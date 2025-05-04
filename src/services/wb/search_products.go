package wb

import (
	"context"
	wbAdapters "core/src/adapters/wb"
	wbClients "core/src/clients/wb"
	"core/src/models"
	"core/src/services"
	"fmt"
	"log"
)

type SearchProductsService struct {
	client                  wbClients.Client
	productsResponseAdapter wbAdapters.ProductsResponseAdapter
	imageURLCreator         ImageURLCreator
}

func NewSearchProductsService(
	client wbClients.Client,
	productsResponseAdapter wbAdapters.ProductsResponseAdapter,
	imageURLCreator ImageURLCreator,
) SearchProductsService {
	return SearchProductsService{
		client:                  client,
		productsResponseAdapter: productsResponseAdapter,
		imageURLCreator:         imageURLCreator,
	}
}

func (s SearchProductsService) SearchProducts(ctx context.Context, request services.SearchProductsServiceRequest) (models.Products, error) {
	// Пока что оставим так
	limit := 10

	getProductsRequest := wbClients.GetProductsRequest{
		Query:       request.Query,
		Sort:        "popular", // popular, priceup
		Limit:       limit,
		Page:        1,
		Currency:    "kzt",
		Language:    "ru",
		Destination: "233", // Значение города Алматы
	}

	// Получаем продукты от WB API
	response, err := s.client.GetProducts(ctx, getProductsRequest)
	if err != nil {
		log.Printf(fmt.Sprintf("[ERROR] can not search, err: %s", err))
		return models.Products{}, err
	}

	// Адаптируем ответ в наши модели
	products := s.productsResponseAdapter.GetProducts(response, limit)

	// Добавим продуктам ссылки на изображения
	for i := range products {
		imageURL, err := s.imageURLCreator.CreateImageURL(products[i].ID)
		if err != nil {
			log.Printf(fmt.Sprintf("[ERROR] can not create image url, productID: %s, err: %s", products[i].ID, err))
			continue
		}

		products[i].ImageURL = imageURL
	}

	return products, nil
}
