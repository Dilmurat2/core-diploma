package adapters

import (
	"core/src/models"
	"fmt"
	"github.com/tidwall/gjson"
	"time"
)

type ProductsResponseAdapter interface {
	GetProducts(response []byte, productsQuantity int) models.Products
}

type productsResponseAdapter struct{}

func NewProductsResponseAdapter() ProductsResponseAdapter {
	return productsResponseAdapter{}
}

func (a productsResponseAdapter) GetProducts(response []byte, productsQuantity int) models.Products {
	products := make(models.Products, 0, productsQuantity)
	productsCounter := productsQuantity

	// Берём массив товаров из data.cards
	rawProducts := gjson.GetBytes(response, "data.cards")
	rawProducts.ForEach(func(key, value gjson.Result) bool {
		product := getProduct(value)
		products = append(products, product)

		productsCounter--
		if productsCounter == 0 {
			return false
		}
		return true
	})

	return products
}

func getProduct(rawProduct gjson.Result) models.Product {
	// Извлекаем нужные поля
	id := rawProduct.Get("id").String()
	name := rawProduct.Get("title").String()

	basicPrice := int(rawProduct.Get("unitPrice").Int())       // "unitPrice"
	currentPrice := int(rawProduct.Get("unitSalePrice").Int()) // "unitSalePrice"

	// Считаем скидку, если есть
	discountPercentage := 0

	// Рейтинг
	rating := rawProduct.Get("rating").Float()

	// Количество отзывов
	reviewsCount := int(rawProduct.Get("reviewsQuantity").Int())

	// При желании добавляем домен Kaspi:
	shopLink := rawProduct.Get("shopLink").String()
	url := fmt.Sprintf("https://www.kaspi.kz/shop%s", shopLink)

	// Первая картинка
	imageURL := rawProduct.Get("previewImages.0.medium").String()

	// Формируем итоговую структуру
	product := models.Product{
		ID:                 id,
		Name:               name,
		URL:                url,
		ImageURL:           imageURL,
		OriginalPrice:      basicPrice,
		CurrentPrice:       currentPrice,
		DiscountPercentage: discountPercentage,
		Rating:             rating,
		ReviewsCount:       reviewsCount,
		DeliveryDate:       calculateDeliveryDate(),
	}

	return product
}

func calculateDeliveryDate() time.Time {
	deliveryTime := time.Now()
	if deliveryTime.Hour() < 13 {
		return deliveryTime
	}
	return deliveryTime.AddDate(0, 0, 1)

}
