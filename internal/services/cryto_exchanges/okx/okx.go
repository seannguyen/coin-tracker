package okx

import (
	"context"
	"fmt"
	"log"

	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/spf13/viper"
)

const exchangeName = "okx"

var client *api.Client

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	apiClient := getApiClient()
	ctx := context.Background()

	// Get account balance from OKX
	balances, err := apiClient.Rest.Account.GetBalance(ctx)
	if err != nil {
		log.Printf("Cannot get balances from OKX: %s", err)
		return nil, err
	}

	log.Println("Get balances from OKX successfully")

	var balanceDataList []*cryto_exchanges.BalanceData

	// Process balance data
	if balances.Code == "0" && len(balances.Data) > 0 {
		for _, accountData := range balances.Data {
			for _, detail := range accountData.Details {
				// Parse available balance
				amount := parseFloat(detail.AvailBal)
				if amount == 0 {
					continue
				}

				// Determine if it's crypto or fiat
				balanceType := cryto_exchanges.Crypto
				if isFiatCurrency(detail.Ccy) {
					balanceType = cryto_exchanges.Fiat
				}

				balance := cryto_exchanges.BalanceData{
					Currency:     detail.Ccy,
					Amount:       amount,
					Type:         balanceType,
					ExchangeName: exchangeName,
				}
				balanceDataList = append(balanceDataList, &balance)
			}
		}
	}

	return balanceDataList, nil
}

func getApiClient() *api.Client {
	if client == nil {
		apiKey := viper.GetString("OKX_API_KEY")
		secretKey := viper.GetString("OKX_API_SECRET")
		passphrase := viper.GetString("OKX_API_PASSPHRASE")

		dest := okex.NormalServer // Use NormalServer for production, or DemoServer for demo trading
		ctx := context.Background()
		client, _ = api.NewClient(ctx, apiKey, secretKey, passphrase, dest)
	}
	return client
}

func parseFloat(s string) float64 {
	var result float64
	_, err := fmt.Sscanf(s, "%f", &result)
	if err != nil {
		return 0
	}
	return result
}

func isFiatCurrency(currency string) bool {
	fiatCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"GBP": true,
		"JPY": true,
		"CNY": true,
		"AUD": true,
		"CAD": true,
		"CHF": true,
		"HKD": true,
		"SGD": true,
		"KRW": true,
		"RUB": true,
		"TRY": true,
		"INR": true,
		"BRL": true,
		"IDR": true,
		"ARS": true,
		"ZAR": true,
	}
	return fiatCurrencies[currency]
}
