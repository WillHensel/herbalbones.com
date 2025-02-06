package square

import (
	"encoding/json"
	"net/url"

	"github.com/google/uuid"
)

type paymentLinkReq struct {
	Order orderReq `json:"order"`
}

type orderReq struct {
	LocationId     string        `json:"location_id"`
	IdempotencyKey string        `json:"idempotency_key"`
	LineItems      []lineItemReq `json:"line_items"`
}

type lineItemReq struct {
	Quantity        string `json:"quantity"`
	CatalogObjectId string `json:"catalog_object_id"`
}

type paymentLinkResp struct {
	PaymentLink PaymentLink `json:"payment_link"`
}

type PaymentLink struct {
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
	Url     string `json:"url"`
}

func (sq *Square) GetSingleItemPaymentLink(itemId string) (string, error) {
	idempotencyKey := uuid.New().String()

	order := paymentLinkReq{
		Order: orderReq{
			LocationId:     sq.locationId,
			IdempotencyKey: idempotencyKey,
			LineItems: []lineItemReq{
				{
					Quantity:        "1",
					CatalogObjectId: itemId,
				},
			},
		},
	}

	reqBody, err := json.Marshal(order)
	if err != nil {
		return "", err
	}

	route, err := url.JoinPath(sq.apiUri, "online-checkout/payment-links")
	if err != nil {
		return "", err
	}

	respBodyBytes, err := sq.makePostRequest(route, reqBody)
	if err != nil {
		return "", err
	}

	respContent := paymentLinkResp{}

	if err := json.Unmarshal(respBodyBytes, &respContent); err != nil {
		return "", err
	}

	return respContent.PaymentLink.Url, nil
}
