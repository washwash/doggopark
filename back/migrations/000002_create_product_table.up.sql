BEGIN;

DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'currency_code') THEN
    CREATE TYPE currency_code AS ENUM ('EUR', 'USD', 'ANG');
  END IF;
END
$$;


CREATE TABLE IF NOT EXISTS price(
  id              serial           PRIMARY KEY,
  amount          NUMERIC          NOT NULL,
  currency_code   currency_code
);


CREATE TABLE IF NOT EXISTS product(
  id          serial   PRIMARY KEY,
  external_id TEXT,
  title       TEXT     NOT NULL,
  description TEXT,
  price_id    INTEGER  REFERENCES price(id)       
);


COMMIT;
