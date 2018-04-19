
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE SEQUENCE snapshots_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE snapshots (
  id INT NOT NULL PRIMARY KEY DEFAULT nextval('snapshots_id_seq'),
  created_at TIMESTAMP NOT NULL
);

CREATE INDEX snapshots_on_created_at ON snapshots USING btree(created_at);
ALTER SEQUENCE snapshots_id_seq OWNED BY snapshots.id;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE snapshots;
DROP SEQUENCE snapshots_id_seq;
