package currency

import (
	"strings"
)

type Service struct {
	validCurrencies map[string]bool
}

func NewService() *Service {
	return &Service{
		validCurrencies: map[string]bool{
			"TRY": true,
			"USD": true,
			"EUR": true,
			"GBP": true,
			"JPY": true,
			"RUB": false,
		},
	}
}

func (s *Service) IsValidCurrency(code string) bool {
	val, ok := s.validCurrencies[strings.ToUpper(code)]
	return ok == val
}

func (s *Service) GetValidCurrencies() []string {
	currencies := make([]string, 0, len(s.validCurrencies))
	for currency, val := range s.validCurrencies {
		if val {
			currencies = append(currencies, currency)
		}
	}
	return currencies
}
