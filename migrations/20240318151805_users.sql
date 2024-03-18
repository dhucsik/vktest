-- +goose Up
-- +goose StatementBegin
CREATE TYPE user_role AS ENUM ('admin', 'user');

CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  role user_role NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TYPE user_role;
-- +goose StatementEnd
