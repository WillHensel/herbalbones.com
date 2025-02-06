package square

import (
	"encoding/json"
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
	IsDeleted     bool                        `json:"is_deleted"`
	IsArchived    bool                        `json:"is_archived"`
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

	route, err := url.JoinPath(sq.apiUri, "catalog/list")
	if err != nil {
		return []CatalogItem{}, err
	}

	route += "?types=ITEM"

	bodyBytes, err := sq.makeGetRequest(route)

	cat := catalogListResp{}

	if err := json.Unmarshal(bodyBytes, &cat); err != nil {
		return []CatalogItem{}, err
	}

	items := []CatalogItem{}

	for _, ob := range cat.Objects {
		if ob.ItemData.IsDeleted || ob.ItemData.IsArchived {
			continue
		}

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
