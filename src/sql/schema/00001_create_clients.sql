-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS clients (
  id SERIAL PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  join_code varchar(255) NOT NULL,
  logo_url varchar(255) NULL,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NULL DEFAULT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS clients;
-- +goose StatementEnd
