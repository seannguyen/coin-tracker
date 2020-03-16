package cmc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/seannguyen/coin-tracker/internal/utilities"

	"github.com/spf13/viper"
)

const (
	apiLimit           = 200
	usdSymbol          = "USD"
	apiEndpoint        = "pro-api.coinmarketcap.com"
	quotePath          = "v1/cryptocurrency/quotes/latest"
	cmcAPIKeyConfigKey = "COIN_MARKET_CAP_API_KEY"
)

var unknownSymbols = []string{"IQX", "MTO", "BAB", "USD"}

// tickerMedia tickers response media
type tickersMedia struct {
	Data map[string]*Ticker `json:"data,omitempty"`
}

// Ticker struct
type Ticker struct {
	ID     int                     `json:"id"`
	Name   string                  `json:"name"`
	Symbol string                  `json:"symbol"`
	Quotes map[string]*TickerQuote `json:"quote"`
}

// TickerQuote struct
type TickerQuote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}

// GetUsdPrices fetch CMC for USD prices
func GetUsdPrices(symbols []string) []float64 {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s/%s", apiEndpoint, quotePath), nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	queriedSymbols := []string{}
	for _, symbol := range symbols {
		if !utilities.Contains(unknownSymbols, symbol) {
			queriedSymbols = append(queriedSymbols, symbol)
		}
	}
	q := url.Values{}
	q.Add("symbol", strings.Join(queriedSymbols, ","))

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", viper.GetString(cmcAPIKeyConfigKey))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		reportErr(err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		reportErr(err)
	}
	if resp.StatusCode != 200 {
		reportErr(string(respBody))
	}

	var body tickersMedia
	err = json.Unmarshal(respBody, &body)
	if err != nil {
		reportErr(err)
	}

	result := []float64{}
	for _, symbol := range symbols {
		if utilities.Contains(unknownSymbols, symbol) {
			result = append(result, 0)
			continue
		}
		symbolData := body.Data[symbol]
		quote := symbolData.Quotes["USD"]
		result = append(result, quote.Price)
	}
	return result
}

func reportErr(err interface{}) {
	panic(fmt.Errorf("fail to fetch price from CMC: %s", err))
}
