# Coin Tracker
Tracker for crypto currencies on various exchanges

## Getting Started
Install DB Migration Tool:
```
go get -v github.com/pressly/goose/cmd/goose
```

## DB Migration
App
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker user=postgres password=password12 sslmode=disable" up
```
Test
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker-test user=postgres password=password12 sslmode=disable" up
```