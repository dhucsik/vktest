package repository

import (
	"context"

	"github.com/dhucsik/vktest/config"
	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
	"github.com/dhucsik/vktest/internal/envs"
	"github.com/dhucsik/vktest/internal/repository/postgre"
)

type ActorRepository interface {
	Create(ctx context.Context, actor *entity.Actor) (int, error)
	Update(ctx context.Context, actor *entity.Actor) error
	GetByID(ctx context.Context, id int) (*entity.Actor, error)
	Delete(ctx context.Context, id int) error
	GetActorsByMovie(ctx context.Context, movieID int) ([]*entity.Actor, error)
	GetActorsByMovies(ctx context.Context, movieIDs []int) (map[int][]*entity.Actor, error)
	SearchActors(ctx context.Context, actor string, limit, offset int) ([]*entity.Actor, error)
	OrderActors(ctx context.Context, limit, offset int) ([]*entity.Actor, error)
}

type MovieRepository interface {
	Create(ctx context.Context, movie *entity.Movie, actorIDs []int) (int, error)
	Update(ctx context.Context, movie *entity.Movie, actorIDs []int) error
	GetByID(ctx context.Context, id int) (*entity.Movie, error)
	Delete(ctx context.Context, id int) error
	Search(ctx context.Context, search string, limit, offset int) ([]*entity.Movie, error)
	GetMoviesByActorIDs(ctx context.Context, actorIDs []int) ([]*entity.Movie, error)
	OrderByRating(ctx context.Context, limit, offset int) ([]*entity.Movie, error)
	OrderByTitle(ctx context.Context, limit, offset int) ([]*entity.Movie, error)
	OrderByReleaseDate(ctx context.Context, limit, offset int) ([]*entity.Movie, error)
	GetMoviesByActors(ctx context.Context, actorID []int) (map[int][]*entity.Movie, error)
	GetMoviesByActorID(ctx context.Context, actorID int) ([]*entity.Movie, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	GetByID(ctx context.Context, id int) (*entity.User, error)
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
}

type Repository struct {
	Actor ActorRepository
	Movie MovieRepository
	User  UserRepository
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	system, err := enum.DatabaseSystemString(cfg.Database.System)
	if err != nil {
		return nil, err
	}

	switch system {
	case enum.DatabaseSystemPostgres:
		return initPostgre(ctx, cfg.Env.Get(envs.PostgresDSN))
	case enum.DatabaseSystemMySQL:
		return nil, nil
	case enum.DatabaseSystemMongo:
		return nil, nil
	default:
		return nil, nil
	}
}

func initPostgre(ctx context.Context, dsn string) (*Repository, error) {
	db, err := postgre.Dial(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &Repository{
		Actor: postgre.NewActorRepo(db),
		Movie: postgre.NewMovieRepo(db),
		User:  postgre.NewUserRepo(db),
	}, nil
}
