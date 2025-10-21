# Coin Tracker
Tracker for crypto currencies on various exchanges

## Supported Exchanges
- Bittrex
- Liquid
- Coinbase
- Bitfinex
- OKX

## Getting Started
Install DB Migration Tool:
```
go get -v github.com/pressly/goose/cmd/goose
```

Install dependencies:
```
dep ensure
```

## Configuration
Copy the sample config file:
```bash
cp configs/config.yml.sample configs/config.yml
```

Edit `configs/config.yml` and add your exchange API credentials:
- For OKX: You'll need `OKX_API_KEY`, `OKX_API_SECRET`, and `OKX_API_PASSPHRASE`
- For other exchanges: Refer to the sample config file

## DB Migration
App
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker user=postgres password=password12 sslmode=disable" up
```
Test
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker-test user=postgres password=password12 sslmode=disable" up
```

