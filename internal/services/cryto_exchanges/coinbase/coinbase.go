package coinbase

import (
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/Zauberstuhl/go-coinbase"
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"errors"
)

const exchangeName = "coinbase"

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	client := getApiClient()
	return getBalances(&client)
}

func getApiClient() coinbase.APIClient {
	return coinbase.APIClient{
		Key: viper.GetString("COINBASE_API_KEY"),
		Secret: viper.GetString("COINBASE_API_SECRET"),
	}
}

func getBalances(client *coinbase.APIClient) ([]*cryto_exchanges.BalanceData, error) {
	accounts, err := client.Accounts()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	if len(accounts.Errors) > 0 {
		errMessages := make([]string, len(accounts.Errors))
		for index, errMessage := range accounts.Errors {
			errMessages[index] = errMessage.Message
		}
		return nil, errors.New(strings.Join(errMessages, ". "))
	}

	balanceByCurrency := make(map[string]float64)
	for _, accountData := range accounts.Data {
		balanceByCurrency[accountData.Currency] += accountData.Native_balance.Amount
	}

	balances := make([]*cryto_exchanges.BalanceData, 0)
	for currency, balance := range balanceByCurrency {
		if balance <= 0 {
			continue
		}
		balanceData := cryto_exchanges.BalanceData{
			Currency: currency,
			Amount: balance,
			ExchangeName: exchangeName,
			Type: cryto_exchanges.Crypto,
		}
		balances = append(balances, &balanceData)
	}

	return balances, nil
}