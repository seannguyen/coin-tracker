package main

import (
	"github.com/spf13/viper"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/seannguyen/coin-tracker/internal/models/snapshot_model"
	"github.com/seannguyen/coin-tracker/internal/services/bittrex"
	"github.com/seannguyen/coin-tracker/internal/models/balance_model"
)

var db *gorm.DB


func main() {
	initialize()
	defer destroy()

	balances, err := bittrex.GetBalances()
	if err != nil {
		log.Panicln(err)
	}

	saveBalancesSnapshot(balances)
}

func initialize() {
	initConfigs()
	initDatabase()
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

func initDatabase() {
	database, err := gorm.Open("postgres", viper.Get("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic(err)
	}
	log.Println("successfully connected to db")
	db = database

	//db.LogMode(true)
}

func destroy() {
	db.Close()
}

func saveBalancesSnapshot(balances []bittrex.BalanceData)  {
	snapshot := snapshot_model.Snapshot{}
	db.Create(&snapshot)

	for _, balance := range balances {
		db.Create(&balance_model.Balance { Currency: balance.Currency, Amount: balance.Amount, Snapshot: snapshot })
	}
}