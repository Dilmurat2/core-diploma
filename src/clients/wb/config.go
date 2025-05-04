package wb

type Config struct {
	SearchProductsURL string // Адрес HTTP метода для поиска продуктов
	ResultSet         string // Код, указывающий для чего именно нужен ответ
	AppType           string // Код клиентского приложения
	Spp               string // Пока не понято зачем, влияние на ответ не заметил
}

func NewConfig() Config {
	return Config{
		SearchProductsURL: "https://search.wb.ru/exactmatch/sng/common/v9/search",
		ResultSet:         "catalog",
		AppType:           "128",
		Spp:               "30",
	}
}
