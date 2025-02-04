package square

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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
	Id      string `json:"id"`
	OrderId string `json:"order_id"`
	Url     string `json:"url"`
}

func (sq *Square) GetSingleItemPaymentLink(itemId string) (string, error) {
	client := &http.Client{}

	idempotencyKey := uuid.New().String()

	order := paymentLinkReq{
		Order: orderReq{
			LocationId:     "LKCMCTJHEAZ72",
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

	reqBodyReader := bytes.NewReader(reqBody)

	route, err := url.JoinPath(sq.apiUri, "online-checkout/payment-links")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", route, reqBodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Add("Square-Version", sq.version)
	req.Header.Add("Authorization", "Bearer "+sq.accessToken)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	respContent := paymentLinkResp{}

	if err := json.Unmarshal(respBodyBytes, &respContent); err != nil {
		return "", err
	}

	return respContent.Url, nil
}
