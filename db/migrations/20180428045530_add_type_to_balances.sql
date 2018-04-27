
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE balances ADD COLUMN type INTEGER NOT NULL DEFAULT 0;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE balances DROP COLUMN type;
