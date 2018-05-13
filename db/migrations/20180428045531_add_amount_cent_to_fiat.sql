-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE fiat_values ADD COLUMN amount_cents BIGINT NOT NULL DEFAULT 0;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE fiat_values DROP COLUMN amount_cents;