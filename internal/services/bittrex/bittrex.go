package bittrex

import (
	"github.com/toorop/go-bittrex"
	"log"
	"github.com/spf13/viper"
)


func GetBalances() ([]bittrex.Balance, error) {
	client := bittrex.New(viper.GetString("BITTREX_API_KEY"), viper.GetString("BITTREX_API_SECRET"))
	balances, err := client.GetBalances()

	if err != nil {
		log.Println("Cannot get balance")
		return nil, err
	}

	log.Print("Get balance successfully")
	return balances, nil
}
