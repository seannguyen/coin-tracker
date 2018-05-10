package cmc

import (
	"github.com/coincircle/go-coinmarketcap"
	"log"
)

const (
	apiLimit  = 200
	usdSymbol = "USD"
)

func GetUsdPrices(symbols []string) []float64 {
	coinsData, err := coinmarketcap.Tickers(&coinmarketcap.TickersOptions{Limit: apiLimit, Convert: usdSymbol})
	if err != nil {
		log.Panic(err)
	}

	prices := make([]float64, len(symbols), len(symbols))
	for index, symbol := range symbols {
		prices[index] = coinsData[symbol].Quotes[usdSymbol].Price
	}

	return prices
}
