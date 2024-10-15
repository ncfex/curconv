package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ncfex/curconv/internal/api"
	"github.com/ncfex/curconv/internal/cli"
	"github.com/ncfex/curconv/internal/converter"
	"github.com/ncfex/curconv/internal/currency"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading environment variables")
	}

	apiKey := os.Getenv("CURRENCY_API_KEY")
	if apiKey == "" {
		log.Fatal("CURRENCY_API_KEY environment variable is not set")
	}

	client := api.NewClient(apiKey)
	currencyService := currency.NewService()
	converterService := converter.NewService(client, currencyService)
	app := cli.NewApp(converterService, currencyService)

	if err := app.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
