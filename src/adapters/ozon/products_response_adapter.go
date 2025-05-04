package ozon

import (
	"core/src/models"
	"fmt"
	"github.com/tidwall/gjson"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var jsonProductPathsMap = map[string]string{
	"name":         `mainState.#(id=="name").atom.textAtom.text`,
	"sku":          "skuId",
	"creditInfo":   `tileImage.leftBottomBadge.text`,
	"deliveryDate": `multiButton.ozonButton.addToCartButtonWithQuantity.text`,
	"priceData":    `mainState.#(atom.type=="priceV2").atom.priceV2`,
	"images":       `tileImage.items`,
	"labels":       `mainState.#(atom.type=="labelList").atom.labelList.items`,
}

type ProductsResponseAdapter interface {
	GetProducts(response string, productsQuantity int) models.Products
}

type productsResponseAdapter struct{}

func NewProductsResponseAdapter() ProductsResponseAdapter {
	return productsResponseAdapter{}
}

func (p productsResponseAdapter) GetProducts(response string, productsQuantity int) models.Products {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(response))
	if err != nil {
		return models.Products{}
	}

	selection := doc.Find("div#state-searchResultsV2-3547909-default-1")
	if selection.Length() == 0 {
		return models.Products{}
	}

	dataState, exists := selection.Attr("data-state")
	if !exists {
		return models.Products{}
	}
	//надо вернуть из начала массива только productsQuantity элементов
	products := p.parseProductItems(dataState)
	if len(products) <= productsQuantity {
		return products
	}

	return products[:productsQuantity]
}

func (p *productsResponseAdapter) parseProductItems(data string) []models.Product {
	var products []models.Product
	items := gjson.Get(data, "items")

	if !items.Exists() {
		return products
	}

	items.ForEach(func(_, item gjson.Result) bool {
		product := p.parseProductItem(item)
		products = append(products, product)
		return true
	})

	return products
}

func (p *productsResponseAdapter) parseProductItem(item gjson.Result) models.Product {
	return models.Product{
		Name:               p.extractName(item),
		ID:                 p.extractSKU(item),
		DeliveryDate:       p.extractDeliveryDate(item),
		URL:                p.generateProductURL(p.extractSKU(item)),
		ImageURL:           p.extractImages(item)[0],
		CurrentPrice:       p.extractCurrentPrice(item),
		OriginalPrice:      p.extractOriginalPrice(item),
		DiscountPercentage: p.extractDiscount(item),
		Rating:             p.extractRating(item),
		ReviewsCount:       p.extractReviewCount(item),
	}
}

func (p *productsResponseAdapter) extractName(item gjson.Result) string {
	return item.Get(jsonProductPathsMap["name"]).String()
}

func (p *productsResponseAdapter) extractSKU(item gjson.Result) string {
	return item.Get(jsonProductPathsMap["sku"]).String()
}

func (p *productsResponseAdapter) extractCreditInfo(item gjson.Result) string {
	return item.Get(jsonProductPathsMap["creditInfo"]).String()
}

func (p *productsResponseAdapter) extractDeliveryDate(item gjson.Result) time.Time {
	deliveryDateStr := item.Get(jsonProductPathsMap["deliveryDate"]).String()
	if deliveryDateStr == "" {
		return time.Time{}
	}

	// Russian month names mapping
	russianMonths := map[string]time.Month{
		"января":   time.January,
		"февраля":  time.February,
		"марта":    time.March,
		"апреля":   time.April,
		"мая":      time.May,
		"июня":     time.June,
		"июля":     time.July,
		"августа":  time.August,
		"сентября": time.September,
		"октября":  time.October,
		"ноября":   time.November,
		"декабря":  time.December,
	}

	deliveryDateStr = strings.TrimSpace(strings.ToLower(deliveryDateStr))

	if deliveryDateStr == "завтра" {
		return time.Now().AddDate(0, 0, 1)
	}
	if deliveryDateStr == "сегодня" {
		return time.Now()
	}

	if deliveryDateStr == "послезавтра" {
		return time.Now().AddDate(0, 0, 2)
	}

	parts := strings.Split(deliveryDateStr, " ")
	if len(parts) != 2 {
		return time.Time{}
	}

	day, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Printf("Failed to parse day: %v", err)
		return time.Time{}
	}

	month, ok := russianMonths[parts[1]]
	if !ok {
		log.Printf("Unknown month: %s", parts[1])
		return time.Time{}
	}

	now := time.Now()
	deliveryDate := time.Date(now.Year(), month, day, 0, 0, 0, 0, time.Local)

	if deliveryDate.Before(now) {
		deliveryDate = time.Date(now.Year()+1, month, day, 0, 0, 0, 0, time.Local)
	}

	return deliveryDate
}

func (p *productsResponseAdapter) generateProductURL(skuID string) string {
	if skuID == "" {
		return ""
	}
	url := fmt.Sprintf("%s%s%s/", "https://ozon.kz", "/product/", skuID)
	return url
}

func (p *productsResponseAdapter) extractCurrentPrice(item gjson.Result) int {
	priceData := item.Get(jsonProductPathsMap["priceData"])
	if !priceData.Exists() {
		return 0
	}
	return p.parsePrice(priceData.Get("price.0.text").String())
}

func (p *productsResponseAdapter) extractOriginalPrice(item gjson.Result) int {
	priceData := item.Get(jsonProductPathsMap["priceData"])
	if !priceData.Exists() {
		return 0
	}
	return p.parsePrice(priceData.Get("price.1.text").String())
}

func (p *productsResponseAdapter) extractDiscount(item gjson.Result) int {
	priceData := item.Get(jsonProductPathsMap["priceData"])
	if !priceData.Exists() {
		return 0
	}
	discountString := priceData.Get("discount").String()

	re := regexp.MustCompile(`-?\d+`)
	match := re.FindString(discountString)
	if match == "" {
		return 0
	}
	number, err := strconv.Atoi(match)
	if err != nil {
		return 0
	}
	return number
}

func (p *productsResponseAdapter) extractRating(item gjson.Result) float64 {
	var ratingStr string

	item.Get("mainState").ForEach(func(_, stateItem gjson.Result) bool {
		if stateItem.Get("atom.type").String() != "labelList" {
			return true
		}

		stateItem.Get("atom.labelList.items").ForEach(func(_, label gjson.Result) bool {
			if strings.Contains(label.Get("icon.image").String(), "star") {
				ratingStr = strings.TrimSpace(label.Get("title").String())
				return false
			}
			return true
		})

		return ratingStr == ""
	})

	if ratingStr == "" {
		return 0
	}

	ratingStr = strings.ReplaceAll(ratingStr, ",", ".")
	rating, err := strconv.ParseFloat(strings.TrimSpace(ratingStr), 64)
	if err != nil {
		log.Printf("Failed to parse rating: %v, err: %v", ratingStr, err)
		return 0
	}

	return rating
}

func (p *productsResponseAdapter) extractReviewCount(item gjson.Result) int {
	var reviewCountStr string

	item.Get("mainState").ForEach(func(_, stateItem gjson.Result) bool {
		if stateItem.Get("atom.type").String() != "labelList" {
			return true
		}

		stateItem.Get("atom.labelList.items").ForEach(func(_, label gjson.Result) bool {
			title := label.Get("title").String()
			if strings.Contains(title, "отзыв") {
				re := regexp.MustCompile(`\d+`)
				matches := re.FindAllString(title, -1)
				reviewCountStr = strings.Join(matches, "")
				return false
			}
			return true
		})

		return reviewCountStr == ""
	})

	if reviewCountStr == "" {
		return 0
	}

	reviewCount, err := strconv.Atoi(strings.TrimSpace(reviewCountStr))
	if err != nil {
		log.Printf("Failed to parse review count: %v", err)
		return 0
	}

	return reviewCount
}

func (p *productsResponseAdapter) extractImages(item gjson.Result) []string {
	var images []string
	item.Get(jsonProductPathsMap["images"]).ForEach(func(_, image gjson.Result) bool {
		if imageURL := image.Get("image.link").String(); imageURL != "" {
			images = append(images, imageURL)
		}
		return true
	})

	return images
}

func (p *productsResponseAdapter) parsePrice(priceText string) int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(priceText, -1)
	if len(matches) == 0 {
		return 0
	}

	price, err := strconv.Atoi(strings.Join(matches, ""))
	if err != nil {
		log.Printf("Failed to parse price: %v", err)
		return 0
	}

	return price
}
