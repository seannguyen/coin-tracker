package coinbase

import (
	"fmt"
	"github.com/Zauberstuhl/go-coinbase"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/spf13/viper"
)

const exchangeName = "coinbase"

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	client := getApiClient()
	return getBalances(&client)
}

func getApiClient() coinbase.APIClient {
	return coinbase.APIClient{
		Key:    viper.GetString("COINBASE_API_KEY"),
		Secret: viper.GetString("COINBASE_API_SECRET"),
	}
}

func getBalances(client *coinbase.APIClient) ([]*cryto_exchanges.BalanceData, error) {
	accounts, err := client.Accounts()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	balanceByCurrency := make(map[string]float64)
	for _, accountData := range accounts.Data {
		balanceByCurrency[accountData.Currency] += accountData.Balance.Amount
	}

	balances := make([]*cryto_exchanges.BalanceData, 0)
	for currency, balance := range balanceByCurrency {
		if balance <= 0 {
			continue
		}
		balanceData := cryto_exchanges.BalanceData{
			Currency:     currency,
			Amount:       balance,
			ExchangeName: exchangeName,
			Type:         cryto_exchanges.Crypto,
		}
		balances = append(balances, &balanceData)
	}

	return balances, nil
}
