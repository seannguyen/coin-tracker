package cmc

import (
	"github.com/miguelmota/go-coinmarketcap"
	"log"
)

const allCoinDataLimit = 200

var currencySymbolMap = make(map[string]string)

func GetUsdPrices(symbols []string) []float64 {
	coinsData := getCoinsDataKeyBySymbol()

	prices := make([]float64, len(symbols), len(symbols))
	for index, symbol := range symbols {
		prices[index] = coinsData[symbol].PriceUSD
	}

	return prices
}

func getCoinsDataKeyBySymbol() map[string]coinmarketcap.Coin {
	coinsData, err := coinmarketcap.GetAllCoinData(allCoinDataLimit)
	if err != nil {
		log.Panic(err)
	}

	dataKeyBySymbol := make(map[string]coinmarketcap.Coin)
	for _, coinData := range coinsData {
		dataKeyBySymbol[coinData.Symbol] = coinData
	}
	return dataKeyBySymbol
}
