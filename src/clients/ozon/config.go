package ozon

type Config struct {
	Site        string            `envconfig:"SITE" default:"ozon.kz"`
	BaseUrl     string            `envconfig:"BASE_URL" default:"https://www.ozon.kz"`
	SearchPath  string            `envconfig:"SEARCH_PATH" default:"/search"`
	ProductPath string            `envconfig:"PRODUCT_PATH" default:"/product"`
	Cookies     map[string]string `envconfig:"COOKIES" default:"adult_user_birthdate:1990-01-01,is_adult_confirmed:true"`
	Selector    string            `envconfig:"SELECTOR" default:"div#state-searchResultsV2-3547909-default-1"`
}

func NewConfig() Config {
	return Config{
		Site:        "ozon.kz",
		BaseUrl:     "https://www.ozon.kz",
		SearchPath:  "/search",
		ProductPath: "/product",
		Cookies:     map[string]string{"adult_user_birthdate": "1990-01-01", "is_adult_confirmed": "true"},
		Selector:    "div#state-searchResultsV2-3547909-default-1",
	}
}
