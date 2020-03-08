package jobs

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bugsnag/bugsnag-go"
	"github.com/gocraft/work"
	_ "github.com/lib/pq" // import for the usage of postgresql
	"github.com/seannguyen/coin-tracker/internal/services/cmc"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/bitfinex"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/bittrex"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/coinbase"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/liquid"
	"github.com/seannguyen/coin-tracker/internal/services/fiat_exchange"
	"github.com/seannguyen/coin-tracker/models"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"gopkg.in/volatiletech/null.v6"
)

// SnapshotBalances pull balances from sources and save them into the DB
func SnapshotBalances(_ *work.Job) error {
	db := initDatabase()
	defer func() {
		err := db.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

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

	saveBalancesSnapshot(db, allBalances)
	return nil
}

func getAllExchanges() []cryto_exchanges.ExchangeInterface {
	return []cryto_exchanges.ExchangeInterface{
		&bittrex.Exchange{},
		&liquid.Exchange{},
		&coinbase.Exchange{},
		&bitfinex.Exchange{},
	}
}

func initDatabase() *sql.DB {
	db, err := sql.Open("postgres", viper.GetString("DB_CONNECTION_STRING"))
	if err != nil {
		log.Panic(err)
	}
	boil.SetDB(db)
	log.Println("Successfully connected to db")
	return db
}

func saveBalancesSnapshot(db *sql.DB, balancesData []*cryto_exchanges.BalanceData) {
	transaction, err := db.Begin()
	if err != nil {
		log.Panic(err)
	}
	defer func() {
		var trxErr error
		if r := recover(); r != nil {
			log.Println("abandon transaction, recovered from ", r)
			trxErr = transaction.Rollback()
			bugsnag.Notify(fmt.Errorf("%v", r))
		} else {
			trxErr = transaction.Commit()
		}
		if trxErr != nil {
			log.Panic(trxErr)
		}
	}()

	snapshot := insertSnapshot(transaction)
	balances := addBalancesToSnapshot(transaction, snapshot, balancesData)
	addUsdValuesToBalances(transaction, balances)
}

func insertSnapshot(transaction *sql.Tx) *models.Snapshot {
	snapshot := models.Snapshot{}
	snapshot.InsertP(transaction)
	log.Println("Successfully create snapshot")
	return &snapshot
}

func addBalancesToSnapshot(transaction *sql.Tx, snapshot *models.Snapshot, balancesData []*cryto_exchanges.BalanceData) []*models.Balance {
	for _, balanceData := range balancesData {
		balance := models.Balance{
			Amount:       balanceData.Amount,
			Currency:     balanceData.Currency,
			ExchangeName: null.StringFrom(balanceData.ExchangeName),
			Type:         balanceData.Type,
		}
		snapshot.AddBalancesP(transaction, true, &balance)
	}
	return snapshot.R.Balances
}

func addUsdValuesToBalances(transaction *sql.Tx, balances []*models.Balance) {
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
	addUsdValueToCryptoBalances(transaction, cryptoBalances)
	addUsdValueToFiatBalances(transaction, fiatBalances)
}

func addUsdValueToCryptoBalances(transaction *sql.Tx, balances []*models.Balance) {
	currencySymbols := make([]string, len(balances), len(balances))
	for index, balance := range balances {
		currencySymbols[index] = balance.Currency
	}

	prices := cmc.GetUsdPrices(currencySymbols)

	for index, price := range prices {
		usdAmountCents := int64(price * balances[index].Amount * 100)
		fiatValue := models.FiatValue{Currency: "USD", AmountCents: usdAmountCents}
		balances[index].AddFiatValuesP(transaction, true, &fiatValue)
	}
}

func addUsdValueToFiatBalances(transaction *sql.Tx, balances []*models.Balance) {
	for _, balance := range balances {
		usdAmount, err := fiat_exchange.ConvertToUsd(balance.Currency, balance.Amount)
		if err != nil {
			panic(err)
		}
		fiatValue := models.FiatValue{Currency: "USD", AmountCents: int64(usdAmount * 100)}
		balance.AddFiatValuesP(transaction, true, &fiatValue)
	}
}
