package wb

import (
	"fmt"
	"strconv"
)

type ImageURLCreator interface {
	CreateImageURL(productID string) (string, error)
}

type imageURLCreator struct {
	imageURLTemplate string
	basketsStorage   BasketsStorage
}

func NewImageURLCreator(
	basketsStorage BasketsStorage,
) ImageURLCreator {
	return imageURLCreator{
		imageURLTemplate: "https://basket-%s.wbbasket.ru/vol%s/part%s/%s/images/c516x688/1.webp",
		basketsStorage:   basketsStorage,
	}
}

func (s imageURLCreator) CreateImageURL(productID string) (string, error) {
	id, err := strconv.Atoi(productID)
	if err != nil {
		return "", err
	}

	basket := s.basketsStorage.GetBasket(id)
	vol, part := s.getVolAndPart(productID)
	url := fmt.Sprintf(s.imageURLTemplate, basket, vol, part, productID)

	return url, nil
}

func (s imageURLCreator) getVolAndPart(productID string) (string, string) {
	vol := ""
	part := ""

	if len(productID) == 8 {
		vol = productID[:3]
		part = productID[:5]
	} else if len(productID) == 9 {
		vol = productID[:4]
		part = productID[:6]
	}

	return vol, part
}
