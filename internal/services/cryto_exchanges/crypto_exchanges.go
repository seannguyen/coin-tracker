package cryto_exchanges

const (
	Crypto = 0
	Fiat   = 1
)

type BalanceData struct {
	Currency string
	Amount   float64
	Type     int
}

type ExchangeInterface interface {
	GetBalances() ([]*BalanceData, error)
}
