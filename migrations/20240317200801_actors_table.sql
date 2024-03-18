-- +goose Up
-- +goose StatementBegin
CREATE TABLE actors (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    gender VARCHAR(25) NOT NULL,
    date_of_birth TIMESTAMP WITH TIME ZONE NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE actors;
-- +goose StatementEnd
