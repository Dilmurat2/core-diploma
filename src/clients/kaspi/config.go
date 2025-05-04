package kaspi

type Config struct {
	SearchProductsURL string // Адрес HTTP метода для поиска продуктов
}

func NewConfig() Config {
	return Config{
		SearchProductsURL: "https://kaspi.kz/yml/product-view/pl/filters",
	}
}
