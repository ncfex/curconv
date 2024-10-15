package cli

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/ncfex/curconv/internal/converter"
	"github.com/ncfex/curconv/internal/currency"
	"github.com/ncfex/curconv/internal/utils"
)

type App struct {
	converter *converter.Service
	currency  *currency.Service
}

func NewApp(converter *converter.Service, currency *currency.Service) *App {
	return &App{
		converter: converter,
		currency:  currency,
	}
}

func (a *App) Run() error {
	var baseCurrency, targetCurrency string
	var amount string

	validCurrencies := a.currency.GetValidCurrencies()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Base Currency").
				Description("Select the base currency").
				Options(a.currencyOptions(validCurrencies)...).
				Value(&baseCurrency),

			huh.NewSelect[string]().
				Title("Target Currency").
				Description("Select the target currency").
				Options(a.currencyOptions(validCurrencies)...).
				Value(&targetCurrency),

			huh.NewInput().
				Title("Amount").
				Description("Enter the amount to convert").
				Value(&amount),
		),
	)

	err := form.Run()
	if err != nil {
		return fmt.Errorf("form input error: %w", err)
	}

	amountFloat, err := utils.StringToFloat(amount)
	if err != nil {
		return fmt.Errorf("form input error: %w", err)
	}

	convertedAmount, err := a.converter.Convert(amountFloat, baseCurrency, targetCurrency)
	if err != nil {
		return fmt.Errorf("conversion error: %w", err)
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amountFloat, baseCurrency, convertedAmount, targetCurrency)

	return nil
}

func (a *App) currencyOptions(currencies []string) []huh.Option[string] {
	options := make([]huh.Option[string], len(currencies))
	for i, currency := range currencies {
		options[i] = huh.NewOption(currency, currency)
	}
	return options
}
