package fiat_exchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const exchangeEndpoint = "https://openexchangerates.org/api/latest.json?app_id=e318a1b91e734ebd9e924f028f6aeef7"

type ratesData struct {
	Timestamp int64              `json:"timestamp"`
	Base      string             `json:"base"`
	Rates     map[string]float64 `json:"rates"`
}

const refreshRateInterval = time.Hour * 10

var ratesInfo ratesData

// ConvertToUsd calculates the correct amount of a source currency to USD
func ConvertToUsd(sourceCurrency string, amount float64) (float64, error) {
	if ratesInfo.Timestamp == 0 || time.Unix(ratesInfo.Timestamp, 0).Add(refreshRateInterval).Before(time.Now()) {
		err := updateRatesInfo()
		if err != nil {
			return 0, err
		}
	}
	rate, ok := ratesInfo.Rates[sourceCurrency]
	if !ok || rate == 0 {
		return 0, errors.New(fmt.Sprintf("currency pair USD-%v doesn't available", sourceCurrency))
	}
	return amount / rate, nil
}

func updateRatesInfo() error {
	res, err := http.Get(exchangeEndpoint)
	defer res.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &ratesInfo)
	if err != nil {
		return err
	}
	return nil
}
