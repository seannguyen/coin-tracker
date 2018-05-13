-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE fiat_values DROP COLUMN amount;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE fiat_values ADD COLUMN amount DECIMAL(32, 16) NOT NULL DEFAULT 0;