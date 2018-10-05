-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE INDEX balances_on_exchange_name ON balances(exchange_name);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP INDEX balances_on_exchange_name;