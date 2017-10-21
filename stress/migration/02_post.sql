-- +migrate Up
CREATE TABLE posts (
  id SERIAL NOT NULL PRIMARY KEY,
  description VARCHAR(256) NOT NULL,
  user_id SERIAL NOT NULL,
  created timestamp NOT NULL DEFAULT NOW()
);
-- +migrate Down
DROP TABLE posts;