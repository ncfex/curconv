package converter

import (
	"fmt"

	"github.com/ncfex/curconv/internal/api"
	"github.com/ncfex/curconv/internal/currency"
)

type Service struct {
	apiClient       *api.Client
	currencyService *currency.Service
}

func NewService(apiClient *api.Client, currencyService *currency.Service) *Service {
	return &Service{
		apiClient:       apiClient,
		currencyService: currencyService,
	}
}

func (s *Service) Convert(amount float64, from, to string) (float64, error) {
	if !s.currencyService.IsValidCurrency(from) || !s.currencyService.IsValidCurrency(to) {
		return 0, fmt.Errorf("invalid currency code")
	}

	rates, err := s.apiClient.GetExchangeRates(from)
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rates: %w", err)
	}

	rate, ok := rates.Rates[to]
	if !ok {
		return 0, fmt.Errorf("exchange rate not found for %s to %s", from, to)
	}

	return amount * rate, nil
}
