package square

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type CatalogItem struct {
	Id          string
	Name        string
	Description string
	Variations  []CatalogItemVariation
	ImageUrls   []string
}

type CatalogItemVariation struct {
	Id    string
	Name  string
	Price int
}

type catalogListResp struct {
	Objects []catalogItemResp `json:"objects"`
}

type catalogItemResp struct {
	Id       string              `json:"id"`
	ItemData catalogItemDataResp `json:"item_data"`
}

type catalogItemDataResp struct {
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	Variations    []catalogItemVariationsResp `json:"variations"`
	EcomImageUrls []string                    `json:"ecom_image_uris"`
}

type catalogItemVariationsResp struct {
	Id                string                `json:"id"`
	ItemVariationData itemVariationDataResp `json:"item_variation_data"`
}

type itemVariationDataResp struct {
	Name        string                 `json:"name"`
	PricingType string                 `json:"pricing_type"`
	PriceMoney  itemVariationPriceResp `json:"price_money"`
}

type itemVariationPriceResp struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func (sq *Square) CatalogList() ([]CatalogItem, error) {
	client := &http.Client{}
	route, err := url.JoinPath(sq.apiUri, "catalog/list")
	if err != nil {
		return []CatalogItem{}, err
	}

	route += "?types=ITEM"

	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		return []CatalogItem{}, err
	}

	req.Header.Add("Square-Version", sq.version)
	req.Header.Add("Authorization", "Bearer "+sq.accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return []CatalogItem{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []CatalogItem{}, err
	}

	cat := catalogListResp{}

	if err := json.Unmarshal(bodyBytes, &cat); err != nil {
		return []CatalogItem{}, err
	}

	items := []CatalogItem{}

	for _, ob := range cat.Objects {
		item := CatalogItem{
			Id:          ob.Id,
			Name:        ob.ItemData.Name,
			Description: ob.ItemData.Description,
			Variations:  []CatalogItemVariation{},
			ImageUrls:   ob.ItemData.EcomImageUrls,
		}

		for _, v := range ob.ItemData.Variations {
			variation := CatalogItemVariation{
				Id:    v.Id,
				Name:  v.ItemVariationData.Name,
				Price: v.ItemVariationData.PriceMoney.Amount,
			}

			item.Variations = append(item.Variations, variation)
		}

		items = append(items, item)
	}

	return items, err
}
