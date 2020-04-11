BEGIN;


CREATE TABLE IF NOT EXISTS account(
  id        serial  PRIMARY KEY,

  username  TEXT    UNIQUE NOT NULL,
  password  TEXT    NOT NULL,
  email     TEXT    UNIQUE NOT NULL
);


COMMIT;
