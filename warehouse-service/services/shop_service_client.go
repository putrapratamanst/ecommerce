package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShopServiceClient struct {
	shopServiceURL string
}

func NewShopServiceClient(shopServiceURL string) *ShopServiceClient {
	return &ShopServiceClient{
		shopServiceURL: shopServiceURL,
	}
}

type ShopDetailResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s *ShopServiceClient) GetShopByID(shopID string) (*ShopDetailResponse, error) {
    url := fmt.Sprintf("%s/shops/%s", s.shopServiceURL, shopID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var shopResponse ShopDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&shopResponse); err != nil {
		return nil, err
	}
	return &shopResponse, nil
}