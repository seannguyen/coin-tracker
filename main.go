package main

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
	"github.com/seannguyen/coin-tracker/internal/services/bittrex"
)

func main() {
	initConfigs()
	balances, err := bittrex.GetBalances()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(balances)
}

func initConfigs() {
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}