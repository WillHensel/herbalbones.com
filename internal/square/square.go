package square

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Square struct {
	accessToken string
	apiUri      string
	version     string
	locationId  string
}

func NewSquare(accessToken string) Square {
	return Square{
		accessToken: accessToken,
		apiUri:      "https://connect.squareup.com/v2",
		version:     "2025-01-23",
		locationId:  "LKCMCTJHEAZ72",
	}
}

func (sq *Square) makeGetRequest(route string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		return []byte{}, err
	}

	sq.addHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bodyBytes, nil
}

func (sq *Square) makePostRequest(route string, body []byte) ([]byte, error) {
	client := &http.Client{}

	reqBodyReader := bytes.NewReader(body)

	req, err := http.NewRequest("POST", route, reqBodyReader)
	if err != nil {
		return []byte{}, err
	}

	sq.addHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []byte{}, errors.New(fmt.Sprintf("request failed\n\n%s", string(bodyBytes)))
	}

	return bodyBytes, nil
}

func (sq *Square) addHeaders(req *http.Request) {
	req.Header.Add("Square-Version", sq.version)
	req.Header.Add("Authorization", "Bearer "+sq.accessToken)
	req.Header.Add("Content-Type", "application/json")
}
