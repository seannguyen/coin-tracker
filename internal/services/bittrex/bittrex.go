package bittrex

import (
	"github.com/spf13/viper"
	"github.com/toorop/go-bittrex"
	"log"
)

type BalanceData struct {
	Currency string
	Amount   float64
}

var client *bittrex.Bittrex

func GetBalances() ([]BalanceData, error) {
	balances, err := apiClient().GetBalances()
	if err != nil {
		log.Printf("Cannot get balances: %s", err)
		return nil, err
	}

	log.Println("Get balances successfully")

	balanceDataList := make([]BalanceData, len(balances))

	for index, balance := range balances {
		amount, _ := balance.Balance.Float64()
		balanceDataList[index] = BalanceData{Currency: balance.Currency, Amount: amount}
	}

	return balanceDataList, nil
}

func apiClient() *bittrex.Bittrex {
	if client == nil {
		client = bittrex.New(viper.GetString("BITTREX_API_KEY"), viper.GetString("BITTREX_API_SECRET"))
	}
	return client
}