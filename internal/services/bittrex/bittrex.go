package bittrex

import (
	"github.com/toorop/go-bittrex"
	"log"
	"github.com/spf13/viper"
	"github.com/seannguyen/coin-tracker/internal/logger"
)

type BalanceData struct {
	Currency string
	Amount float64
}

func GetBalances() ([]BalanceData, error) {
	client := bittrex.New(viper.GetString("BITTREX_API_KEY"), viper.GetString("BITTREX_API_SECRET"))
	balances, err := client.GetBalances()

	if err != nil {
		log.Println("Cannot get balances")
		return nil, err
	}

	logger.Info("Get balances successfully")

	balanceDataList := make([]BalanceData, len(balances))

	for index, balance := range balances  {
		amount, _ := balance.Balance.Float64()
		balanceDataList[index] = BalanceData{ Currency: balance.Currency, Amount: amount }
	}

	return balanceDataList, nil
}
