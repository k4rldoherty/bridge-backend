-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
  id SERIAL PRIMARY KEY,
  role varchar(255) NOT NULL
);

-- +goose StatementBegin
INSERT INTO roles (role) VALUES ('sys');
INSERT INTO roles (role) VALUES ('superuser');
INSERT INTO roles (role) VALUES ('admin');
INSERT INTO roles (role) VALUES ('user');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS roles;
-- +goose StatementEnd
