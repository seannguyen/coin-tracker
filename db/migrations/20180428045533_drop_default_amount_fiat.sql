-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE fiat_values ALTER COLUMN amount_cents DROP DEFAULT;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE fiat_values ALTER COLUMN amount_cents SET DEFAULT 0;