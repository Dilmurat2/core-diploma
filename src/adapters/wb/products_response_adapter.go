package wb

import (
	"core/src/models"
	"fmt"
	"github.com/tidwall/gjson"
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

	rawProducts := gjson.Get(string(response), "data.products")
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
	// Получаем имя товара
	name := getProductName(rawProduct)

	// Получаем цены товара
	basicPrice, totalPrice := getOriginalAndCurrentPrices(rawProduct)

	// Получаем скидку товара
	discountPercentage := getDiscount(basicPrice, totalPrice)

	// Получаем рейтинг товара
	reviewRating := getReviewRating(rawProduct)

	// Получаем количество отзывов
	feedbacksQuantity := getFeedbacksQuantity(rawProduct)

	// Получаем артикул
	id := getID(rawProduct)

	// Получим URL продукта
	url := getProductURL(id)

	product := models.Product{
		ID:                 id,
		Name:               name,
		URL:                url,
		OriginalPrice:      basicPrice / 100,
		CurrentPrice:       totalPrice / 100,
		DiscountPercentage: discountPercentage,
		Rating:             reviewRating,
		ReviewsCount:       feedbacksQuantity,
	}

	return product
}

func getID(rawProduct gjson.Result) string {
	return rawProduct.Get("id").String()
}

func getFeedbacksQuantity(rawProduct gjson.Result) int {
	return int(rawProduct.Get("feedbacks").Int())
}

func getReviewRating(rawProduct gjson.Result) float64 {
	return rawProduct.Get("reviewRating").Float()
}

func getOriginalAndCurrentPrices(rawProduct gjson.Result) (int, int) {
	rawSizes := rawProduct.Get("sizes")
	originalPrice := 0
	currentPrice := 0
	if rawSizes.Exists() {
		firstSize := rawSizes.Array()[0]
		originalPrice = int(firstSize.Get("price.basic").Int())
		currentPrice = int(firstSize.Get("price.total").Int())
	}

	return originalPrice, currentPrice
}

func getProductName(rawProduct gjson.Result) string {
	return rawProduct.Get("name").String()
}

func getDiscount(basicPrice, totalPrice int) int {
	if totalPrice < basicPrice {
		return int((float64(basicPrice-totalPrice) / float64(basicPrice)) * 100)
	}

	return 0
}

func getProductURL(id string) string {
	return fmt.Sprintf("https://www.wildberries.ru/catalog/%s/detail.aspx", id)
}
