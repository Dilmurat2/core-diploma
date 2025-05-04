package ozon

import (
	ozonAdapter "core/src/adapters/ozon"
	ozonClient "core/src/clients/ozon"
	ozonServices "core/src/services/ozon"
	"fmt"
)

func InitOzonSearchProductsService() (ozonServices.SearchProductsService, error) {
	clientConfig := ozonClient.NewConfig()
	ozonClient, err := ozonClient.NewClient(clientConfig)
	if err != nil {
		return ozonServices.SearchProductsService{}, fmt.Errorf("failed to initialize ozon client: %w", err)
	}

	ozonSearchProductsResponseAdapter := ozonAdapter.NewProductsResponseAdapter()

	ozonSearchProductsService := ozonServices.NewSearchProductsService(ozonClient, ozonSearchProductsResponseAdapter)

	return ozonSearchProductsService, nil
}
