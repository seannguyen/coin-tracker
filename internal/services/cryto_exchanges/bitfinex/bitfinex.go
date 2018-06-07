package bitfinex

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/rhymond/go-money"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/spf13/viper"
	"strconv"
	"strings"
)

const (
	exchangeName = "bitfinex"
)

var externalCurrenciesMap = map[string]string{
	"IOT": "MIOTA",
}

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	client := bitfinex.NewClient().Auth(viper.GetString("BITFINEX_API_KEY"), viper.GetString("BITFINEX_API_SECRET"))
	balances, err := client.Balances.All()

	if err != nil {
		return nil, err
	}

	var balancesData []*cryto_exchanges.BalanceData
	for _, balance := range balances {
		amount, err := strconv.ParseFloat(balance.Amount, 64)
		if err != nil {
			return nil, err
		}
		if amount <= 0 {
			continue
		}
		currency := strings.ToUpper(balance.Currency)
		balancesData = append(balancesData, &cryto_exchanges.BalanceData{
			Type:         getCurrencyType(currency),
			Amount:       amount,
			ExchangeName: exchangeName,
			Currency:     getExternalCurrencySymbol(currency),
		})
	}
	return balancesData, nil
}

func getCurrencyType(currency string) int {
	currencyRecord := money.GetCurrency(currency)
	if currencyRecord != nil {
		return cryto_exchanges.Fiat
	}
	return cryto_exchanges.Crypto
}

func getExternalCurrencySymbol(internalSymbol string) string {
	externalCurrency, ok := externalCurrenciesMap[internalSymbol]
	if ok {
		return externalCurrency
	}
	return internalSymbol
}
