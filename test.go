package main

import (
	"log"
	"github.com/spf13/viper"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/bittrex"
)

func main()  {
	initConfigs()
	exchange := bittrex.Exchange{}
	balances, _ := exchange.GetBalances()
	dataJson,_ :=json.MarshalIndent(balances, "", "	")
	fmt.Println(string(dataJson))
}
func initConfigs() {
	log.Println("Initializing configs")
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}
}