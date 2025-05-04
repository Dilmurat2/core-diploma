package models

import "time"

type Product struct {
	ID                 string    `json:"id"`                  // Артикул или другой идентификатор
	Name               string    `json:"name"`                // Наименование продукта
	URL                string    `json:"url"`                 // Ссылка на страницу продукта (pdp)
	ImageURL           string    `json:"image_url"`           // Ссылка на изображение продукта
	CurrentPrice       int       `json:"current_price"`       // Текущая цена, по которой продают
	OriginalPrice      int       `json:"original_price"`      // Базовая цена, от которой считается скидка
	DiscountPercentage int       `json:"discount_percentage"` // Скидка в процентах
	ReviewsCount       int       `json:"reviews_count"`       // Количество отзывов
	Rating             float64   `json:"rating"`              // Рейтинг продукта
	DeliveryDate       time.Time `json:"delivery_date"`       // Дата доставки
}
