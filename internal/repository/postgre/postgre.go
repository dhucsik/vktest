package postgre

import (
	"context"

	"github.com/dhucsik/vktest/internal/repository/postgre/actor"
	"github.com/dhucsik/vktest/internal/repository/postgre/movie"
	"github.com/dhucsik/vktest/internal/repository/postgre/user"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Dial(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewActorRepo(db *pgxpool.Pool) *actor.Repository {
	return actor.NewRepository(db)
}

func NewMovieRepo(db *pgxpool.Pool) *movie.Repository {
	return movie.NewRepository(db)
}

func NewUserRepo(db *pgxpool.Pool) *user.Repository {
	return user.NewRepository(db)
}
