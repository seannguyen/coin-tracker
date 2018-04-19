package jobs

import (
	"github.com/seannguyen/coin-tracker/internal/services/bittrex"
	"github.com/seannguyen/coin-tracker/internal/models/snapshot_model"
	"github.com/seannguyen/coin-tracker/internal/models/balance_model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"github.com/gocraft/work"
)

var db *gorm.DB

func init() {
	initDatabase()
}

func SnapshotBalances(_ *work.Job) error {
	balances, err := bittrex.GetBalances()
	if err != nil {
		return err
	}

	saveBalancesSnapshot(balances)
	return nil
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