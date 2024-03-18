-- +goose Up
-- +goose StatementBegin
CREATE TABLE movies (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    release_date DATE NOT NULL,
    rating NUMERIC(3, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE movies_actors (
    movie_id INT REFERENCES movies(id),
    actor_id INT REFERENCES actors(id),
    PRIMARY KEY (movie_id, actor_id)
);
-- +goose StatementEnd

-- +goose Down

-- +goose StatementBegin
DROP TABLE movies_actors;
-- +goose StatementEnd

-- +goose StatementBegin
DROP TABLE movies;
-- +goose StatementEnd
