
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE SEQUENCE balances_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE balances (
  id INT NOT NULL PRIMARY KEY DEFAULT nextval('balances_id_seq'),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  snapshot_id INT NOT NULL REFERENCES snapshots(id),
  currency VARCHAR(10) NOT NULL,
  amount DECIMAL(32, 16)  NOT NULL
);

CREATE INDEX balances_on_snapshot_id ON balances USING btree(snapshot_id);
CREATE INDEX balances_on_currency ON balances USING btree(currency);
ALTER SEQUENCE balances_id_seq OWNED BY balances.id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE balances;