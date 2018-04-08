
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE SEQUENCE fiat_values_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE fiat_values (
  id INT NOT NULL PRIMARY KEY DEFAULT nextval('fiat_values_id_seq'),
  balance_id INT NOT NULL REFERENCES balances(id),
  currency VARCHAR(10) NOT NULL,
  amount DECIMAL(32, 16)  NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE fiat_values;
DROP SEQUENCE fiat_values_id_seq;