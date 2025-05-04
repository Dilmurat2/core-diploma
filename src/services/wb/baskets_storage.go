package wb

type BasketsStorage interface {
	GetBasket(id int) string
}

type basketsStorage struct {
	storage map[[2]int]string
}

func NewBasketsStorage() BasketsStorage {
	storage := make(map[[2]int]string)

	storage[[2]int{10000000, 14399999}] = "01"
	storage[[2]int{14400000, 28799999}] = "02"
	storage[[2]int{28800000, 43199999}] = "03"
	storage[[2]int{43200000, 71999999}] = "04"
	storage[[2]int{72000000, 100799999}] = "05"
	storage[[2]int{100800000, 106199999}] = "06"
	storage[[2]int{106200000, 111599999}] = "07"
	storage[[2]int{111600000, 116999999}] = "08"
	storage[[2]int{117000000, 131399999}] = "09"
	storage[[2]int{131400000, 160199999}] = "10"
	storage[[2]int{160200000, 165599999}] = "11"
	storage[[2]int{165600000, 191999999}] = "12"
	storage[[2]int{192000000, 204599999}] = "13"
	storage[[2]int{204600000, 218999999}] = "14"
	storage[[2]int{219000000, 240599999}] = "15"
	storage[[2]int{240600000, 262199999}] = "16"
	storage[[2]int{262200000, 283799999}] = "17"
	storage[[2]int{283800000, 305399999}] = "18"
	storage[[2]int{305400000, 316123969}] = "19"

	return basketsStorage{storage: storage}
}

func (s basketsStorage) GetBasket(id int) string {
	for _range, basketNumber := range s.storage {
		if s.isBetween(id, _range[0], _range[1]) {
			return basketNumber
		}
	}

	return ""
}

func (s basketsStorage) isBetween(num, lower, upper int) bool {
	return num > lower && num < upper
}
