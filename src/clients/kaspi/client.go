package kaspi

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
	restyClient := resty.New().
		SetDebug(true)

	params := map[string]string{
		"c":    "750000000",
		"text": request.Query,
		"sort": request.Sort,
		"page": fmt.Sprintf("%d", 0),
	}

	headers := map[string]string{
		"Referer":    "https://kaspi.kz/shop/search/?text=iphone%2011&hint_chips_click=false",
		"Accept":     "application/json, text/*",
		"X-Ks-City":  "750000000",
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0",
	}

	response, err := restyClient.R().
		SetContext(ctx).
		SetQueryParams(params).
		SetHeaders(headers).
		Get(c.config.SearchProductsURL)
	if err != nil {
		return nil, err
	}
	if response.IsError() {
		return nil, errors.New(response.String())
	}

	body := response.Body()
	return body, nil

}
