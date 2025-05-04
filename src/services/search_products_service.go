package services

import (
	"context"
	"core/src/models"
)

// SearchProductsService - все сервисы парсеры должны имплементировать этот интерфейс для работы с агрегатором
type SearchProductsService interface {
	SearchProducts(ctx context.Context, request SearchProductsServiceRequest) (models.Products, error)
}
