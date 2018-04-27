package bittrex

import (
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/spf13/viper"
	"github.com/toorop/go-bittrex"
	"log"
)

var client *bittrex.Bittrex

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	balances, err := apiClient().GetBalances()
	if err != nil {
		log.Printf("Cannot get balances: %s", err)
		return nil, err
	}

	log.Println("Get balances successfully")

	var balanceDataList []*cryto_exchanges.BalanceData

	for _, balance := range balances {
		amount, _ := balance.Balance.Float64()
		if amount == 0 {
			continue
		}
		balance := cryto_exchanges.BalanceData{
			Currency: balance.Currency,
			Amount: amount,
			Type: cryto_exchanges.Crypto,
			ExchangeName: "bittrex",
		}
		balanceDataList = append(balanceDataList, &balance)
	}

	return balanceDataList, nil
}

func apiClient() *bittrex.Bittrex {
	if client == nil {
		client = bittrex.New(viper.GetString("BITTREX_API_KEY"), viper.GetString("BITTREX_API_SECRET"))
	}
	return client
}
