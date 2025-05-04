package wb

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client interface {
	GetProducts(context.Context, GetProductsRequest) ([]byte, error)
}

type client struct {
	config Config
}

func NewClient(config Config) Client {
	return client{config: config}
}

func (c client) GetProducts(ctx context.Context, request GetProductsRequest) ([]byte, error) {
	restyClient := resty.New().SetDebug(true)

	params := map[string]string{
		"query":     request.Query,
		"resultset": c.config.ResultSet,
		"limit":     fmt.Sprintf("%d", request.Limit),
		"sort":      request.Sort,
		"page":      fmt.Sprintf("%d", request.Page),
		"appType":   c.config.AppType,
		"curr":      request.Currency,
		"lang":      request.Language,
		"dest":      request.Destination,
		"spp":       c.config.Spp,
	}

	response, err := restyClient.R().SetContext(ctx).SetQueryParams(params).Get(c.config.SearchProductsURL)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	body := response.Body()

	return body, nil
}
