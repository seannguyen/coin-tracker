package liquid

import (
	"encoding/json"
	"fmt"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges"
	"github.com/seannguyen/coin-tracker/internal/services/cryto_exchanges/liquid/jwt"
	"io/ioutil"
	"net/http"
)

type account struct {
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance,string"`
}

type accountsContainer struct {
	CryptoAccounts []*account `json:"crypto_accounts"`
	FiatAccounts   []*account `json:"fiat_accounts"`
}

const (
	accountApiEndPoint = "https://api.quoine.com/accounts"
)

type Exchange struct{}

func (*Exchange) GetBalances() ([]*cryto_exchanges.BalanceData, error) {
	rawAccountData, err := getRawAccountData()
	if err != nil {
		return nil, err
	}

	var data accountsContainer
	err = json.Unmarshal(rawAccountData, &data)
	if err != nil {
		return nil, err
	}

	cryptoBalances := constructBalancesData(data.CryptoAccounts, cryto_exchanges.Crypto)
	fiatBalances := constructBalancesData(data.FiatAccounts, cryto_exchanges.Fiat)
	balances := append(cryptoBalances, fiatBalances...)
	return balances, nil
}

func constructBalancesData(accounts []*account, balanceType int) []*cryto_exchanges.BalanceData {
	var balances []*cryto_exchanges.BalanceData
	for _, account := range accounts {
		if account.Balance == 0 {
			continue
		}
		balance := cryto_exchanges.BalanceData{
			Currency:     account.Currency,
			Amount:       account.Balance,
			Type:         balanceType,
			ExchangeName: "liquid",
		}
		balances = append(balances, &balance)
	}
	return balances
}

func getRawAccountData() ([]byte, error) {
	client := http.Client{}
	request, err := constructRequest()
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"fail to get balance data, status code %d, message '%s'", response.StatusCode, string(body))
	}

	return body, nil
}

func constructRequest() (*http.Request, error) {
	request, err := http.NewRequest("GET", accountApiEndPoint, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("X-Quoine-API-Version", "2")
	request.Header.Add("X-Quoine-Auth", jwt.Create())
	request.Header.Add("Content-Type", "application/json")
	return request, nil
}
