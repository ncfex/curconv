package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "https://api.exchangerate-api.com/v4/latest/"
const storePath = "store"

type Client struct {
	apiKey     string
	httpClient *http.Client
}

type ExchangeRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetExchangeRates(baseCurrency string) (*ExchangeRates, error) {
	url := fmt.Sprintf("%s%s?api_key=%s", baseURL, baseCurrency, c.apiKey)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	var rates ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&rates); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	jsonData, err := json.Marshal(rates)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	err = os.MkdirAll(storePath, 0700)
	if err != nil {
		return nil, fmt.Errorf("failed to create store directory: %w", err)
	}

	filename := fmt.Sprintf(storePath+"/%v_%s.json", strings.ReplaceAll(rates.Date, ":", "-"), baseCurrency)
	err = os.WriteFile(filename, jsonData, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &rates, nil
}
