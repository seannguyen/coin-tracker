# Coin Tracker

## DB Migration
App
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker user=seannguyen password=password12" up
```
Test
```bash
goose -dir db/migrations/ postgres "host=localhost database=coin-tracker-test user=seannguyen password=password12" up
```
