package wb

import (
	wbAdapters "core/src/adapters/wb"
	wbClients "core/src/clients/wb"
	"core/src/services"
	wbServices "core/src/services/wb"
)

func InitWBSearchProductsService() services.SearchProductsService {
	// init client
	clientConfig := wbClients.NewConfig()
	client := wbClients.NewClient(clientConfig)

	// init adapter
	productsResponseAdapter := wbAdapters.NewProductsResponseAdapter()

	// init support services
	basketStorage := wbServices.NewBasketsStorage()
	imageURLCreator := wbServices.NewImageURLCreator(basketStorage)

	// init search products service
	searchProductsService := wbServices.NewSearchProductsService(
		client,
		productsResponseAdapter,
		imageURLCreator,
	)

	return searchProductsService
}
