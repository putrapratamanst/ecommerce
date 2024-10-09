package services

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProductServiceClient struct {
	productServiceURL string
}

func NewProductServiceClient(productServiceURL string) *ProductServiceClient {
	return &ProductServiceClient{
		productServiceURL: productServiceURL,
	}
}

type ProductDetailResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		ID       uint    `json:"ID"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity string  `json:"quantity"`
	} `json:"data"`
}

func (s *ProductServiceClient) GetProductByID(productID uint) (*ProductDetailResponse, error) {
	url := fmt.Sprintf("%s/products/%d", s.productServiceURL, productID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var productResponse ProductDetailResponse
	if err := json.NewDecoder(resp.Body).Decode(&productResponse); err != nil {
		return nil, err
	}
	return &productResponse, nil
}
