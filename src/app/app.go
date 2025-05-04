package app

import (
	"core/src/app/kaspi"
	"core/src/app/ozon"
	"core/src/app/wb"
	searchProductsResource "core/src/resources/http/v1/core/search_products"
	"core/src/services"
	searchProductsAggregator "core/src/services/core/search_products_aggregator"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func InitServer() (chi.Router, error) {
	// init wb search products service
	wbSearchProductsService := wb.InitWBSearchProductsService()

	kaspiSearchProductsService := kaspi.InitKaspiSearchProductsService()

	ozonSearchProductsService, err := ozon.InitOzonSearchProductsService()
	if err != nil {
		return nil, err
	}

	// define search products services map
	searchProductsServices := make(map[searchProductsAggregator.SourceType]services.SearchProductsService)
	searchProductsServices[searchProductsAggregator.WB] = wbSearchProductsService
	searchProductsServices[searchProductsAggregator.KASPI] = kaspiSearchProductsService
	searchProductsServices[searchProductsAggregator.OZON] = ozonSearchProductsService

	// init search products aggregator
	aggregator := searchProductsAggregator.NewAggregator(searchProductsServices)

	// init router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Heartbeat("/health"))
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	}))

	// init resources
	resource := searchProductsResource.NewResource(aggregator)
	router.Get("/api/v1/products/search", resource.SearchProducts)

	return router, nil
}
