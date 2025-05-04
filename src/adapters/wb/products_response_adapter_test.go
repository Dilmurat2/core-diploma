package wb

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"wb_parser/models"
)

// Проверка длины массива продуктов от адаптера
func Test_GetProductsFromResponse_case_1(t *testing.T) {
	adapter := NewProductsResponseAdapter()
	response, _ := os.ReadFile("./test_data/testcase_1.json")
	products := adapter.GetProducts(response, 10)
	assert.Equal(t, 10, len(products))
}

// Проверка заполненности полей в модели продукта models.Product
func Test_GetProductsFromResponse_case_2(t *testing.T) {
	adapter := NewProductsResponseAdapter()
	response, _ := os.ReadFile("./test_data/testcase_1.json")
	products := adapter.GetProducts(response, 4)
	assert.Equal(t, 4, len(products))

	product := products[3]
	expectedProduct := models.Product{
		ID:                 "274592574",
		Name:               "Напиток овсяный шоколадный 0,2л, 14 шт",
		ProductURL:         "https://www.wildberries.ru/catalog/274592574/detail.aspx",
		ProductImageURL:    "https://basket-17.wbbasket.ru/vol2765/part276595/276595931/images/c516x688/1.webp",
		CurrentPrice:       6350,
		OriginalPrice:      6810,
		DiscountPercentage: 6,
		FeedbacksQuantity:  1,
		ReviewRating:       5,
	}
	assert.Equal(t, expectedProduct, product)
}
