package jobs

import (
	"github.com/gocraft/work"
	"github.com/seannguyen/coin-tracker/internal/services/bittrex"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/seannguyen/coin-tracker/models"
	"gopkg.in/volatiletech/null.v6"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/seannguyen/coin-tracker/internal/services/cmc"
)

var db *sql.DB

func SnapshotBalances(_ *work.Job) error {
	initDatabase()
	balances, err := bittrex.GetBalances()
	if err != nil {
		return err
	}
	log.Println("Successfully fetching bittrex balances")

	saveBalancesSnapshot(balances)
	return nil
}

func initDatabase() {
	var err error
	db, err = sql.Open("postgres", viper.GetString("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic(err)
	}
	boil.SetDB(db)
	log.Println("Successfully connected to db")
}

func saveBalancesSnapshot(balancesData []bittrex.BalanceData) {
	transaction, err := db.Begin()
	if err != nil { log.Panic(err) }
	snapshot := insertSnapshot()
	balances := addBalancesToSnapshot(snapshot, balancesData)
	addFiatValuesToBalances(balances)
	transaction.Commit()
}

func insertSnapshot() *models.Snapshot {
	snapshot := models.Snapshot{}
	snapshot.InsertGP()
	log.Println("Successfully create snapshot")
	return &snapshot
}

func addBalancesToSnapshot(snapshot *models.Snapshot, balancesData []bittrex.BalanceData) []*models.Balance {
	for _, balanceData := range balancesData {
		balance := models.Balance{
			Amount: balanceData.Amount,
			Currency: balanceData.Currency,
			ExchangeName: null.StringFrom("bittrex"),
		}
		snapshot.AddBalancesGP(true, &balance)
	}
	return snapshot.R.Balances
}

func addFiatValuesToBalances(balances []*models.Balance) {
	currencySymbols := make([]string, len(balances), len(balances))
	for index, balance := range balances {
		currencySymbols[index] = balance.Currency
	}

	prices := cmc.GetUsdPrices(currencySymbols)

	for index, price := range prices {
		amount := price * balances[index].Amount
		fiatValue := models.FiatValue{ Currency: "USD",	Amount: amount }
		balances[index].AddFiatValuesGP(true, &fiatValue)
	}
}