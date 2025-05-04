package kaspi

import (
	kaspiAdapter "core/src/adapters/kaspi"
	kaspiClients "core/src/clients/kaspi"
	kaspiServices "core/src/services/kaspi"
)

func InitKaspiSearchProductsService() kaspiServices.SearchProductsService {
	clientConfig := kaspiClients.NewConfig()

	kaspiClient := kaspiClients.NewClient(clientConfig)

	kaspiProductsResponseAdapter := kaspiAdapter.NewProductsResponseAdapter()

	kaspiSearchProductsService := kaspiServices.NewSearchProductsService(
		kaspiClient,
		kaspiProductsResponseAdapter)

	return kaspiSearchProductsService
}
