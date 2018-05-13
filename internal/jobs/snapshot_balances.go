package jobs

import (
	"database/sql"
	"github.com/gocraft/work"
	_ "github.com/lib/pq"
	"github.com/seannguyen/coin-tracker/internal/services/cmc"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/bittrex"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/quoinex"
	"github.com/seannguyen/coin-tracker/internal/services/fiat_exchange"
	"github.com/seannguyen/coin-tracker/models"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"gopkg.in/volatiletech/null.v6"
	"log"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/coinbase"
)

var db *sql.DB

func SnapshotBalances(_ *work.Job) error {
	initDatabase()
	defer db.Close()

	var allBalances []*cryto_exchanges.BalanceData
	exchanges := getAllExchanges()
	for _, exchange := range exchanges {
		exchangeBalances, err := exchange.GetBalances()
		if err != nil {
			return err
		}
		allBalances = append(allBalances, exchangeBalances...)
	}

	log.Println("Successfully fetching balances")

	saveBalancesSnapshot(allBalances)
	return nil
}

func getAllExchanges() []cryto_exchanges.ExchangeInterface {
	return []cryto_exchanges.ExchangeInterface{
		&bittrex.Exchange{},
		&quoinex.Exchange{},
		&coinbase.Exchange{},
	}
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

func saveBalancesSnapshot(balancesData []*cryto_exchanges.BalanceData) {
	transaction, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	snapshot := insertSnapshot()
	balances := addBalancesToSnapshot(snapshot, balancesData)
	addUsdValuesToBalances(balances)
	transaction.Commit()
}

func insertSnapshot() *models.Snapshot {
	snapshot := models.Snapshot{}
	snapshot.InsertGP()
	log.Println("Successfully create snapshot")
	return &snapshot
}

func addBalancesToSnapshot(snapshot *models.Snapshot, balancesData []*cryto_exchanges.BalanceData) []*models.Balance {
	for _, balanceData := range balancesData {
		balance := models.Balance{
			Amount:       balanceData.Amount,
			Currency:     balanceData.Currency,
			ExchangeName: null.StringFrom(balanceData.ExchangeName),
			Type:         balanceData.Type,
		}
		snapshot.AddBalancesGP(true, &balance)
	}
	return snapshot.R.Balances
}

func addUsdValuesToBalances(balances []*models.Balance) {
	var cryptoBalances, fiatBalances []*models.Balance
	for _, balance := range balances {
		switch balance.Type {
		case cryto_exchanges.Crypto:
			cryptoBalances = append(cryptoBalances, balance)
		case cryto_exchanges.Fiat:
			fiatBalances = append(fiatBalances, balance)
		default:
			panic("balance is missing type")
		}
	}
	addUsdValueToCryptoBalances(cryptoBalances)
	addUsdValueToFiatBalances(fiatBalances)
}

func addUsdValueToCryptoBalances(balances []*models.Balance) {
	currencySymbols := make([]string, len(balances), len(balances))
	for index, balance := range balances {
		currencySymbols[index] = balance.Currency
	}

	prices := cmc.GetUsdPrices(currencySymbols)

	for index, price := range prices {
		amount := price * balances[index].Amount
		fiatValue := models.FiatValue{Currency: "USD", Amount: amount}
		balances[index].AddFiatValuesGP(true, &fiatValue)
	}
}

func addUsdValueToFiatBalances(balances []*models.Balance) {
	for _, balance := range balances {
		usdAmount, err := fiat_exchange.ConvertToUsd(balance.Currency, balance.Amount)
		if err != nil {
			panic(err)
		}
		fiatValue := models.FiatValue{Currency: "USD", Amount: usdAmount}
		balance.AddFiatValuesGP(true, &fiatValue)
	}
}
