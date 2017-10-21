-- +migrate Up
CREATE TABLE users (
  id SERIAL NOT NULL PRIMARY KEY,
  nickname VARCHAR(256),
  email VARCHAR(256) NOT NULL,
  password VARCHAR(256) NOT NULL,
  created timestamp NOT NULL DEFAULT NOW()
);
-- +migrate Down
DROP TABLE users;