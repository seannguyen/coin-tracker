package main

import (
	"fmt"
	"github.com/toorop/go-bittrex"
)

const (
	API_KEY    = ""
	API_SECRET = ""
)

func main() {
	bittrex := bittrex.New(API_KEY, API_SECRET)
	balances, err := bittrex.GetBalances()
	fmt.Println(err, balances)
}