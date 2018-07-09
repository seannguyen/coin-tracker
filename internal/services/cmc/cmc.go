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
		for _, coinData := range coinsData {
			if coinData.Symbol == symbol {
				prices[index] = coinData.Quotes[usdSymbol].Price
				break
			}
		}
	}

	return prices
}
