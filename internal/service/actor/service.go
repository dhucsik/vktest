package actor

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/repository"
	"github.com/samber/lo"
)

type Service interface {
	CreateActor(ctx context.Context, actor *entity.Actor) (int, error)
	UpdateActor(ctx context.Context, id int, params *entity.UpdateActorParams) error
	DeleteActor(ctx context.Context, id int) error
	OrderActors(ctx context.Context, limit, offset int) ([]*entity.ActorWithMovies, error)
	GetActor(ctx context.Context, id int) (*entity.ActorWithMovies, error)
}

type service struct {
	actorRepo  repository.ActorRepository
	moviesRepo repository.MovieRepository
}

func NewService(
	actorRepo repository.ActorRepository,
	moviesRepo repository.MovieRepository,
) Service {
	return &service{
		actorRepo:  actorRepo,
		moviesRepo: moviesRepo,
	}
}

func (s *service) CreateActor(ctx context.Context, actor *entity.Actor) (int, error) {
	return s.actorRepo.Create(ctx, actor)
}

func (s *service) UpdateActor(ctx context.Context, id int, params *entity.UpdateActorParams) error {
	actr, err := s.actorRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	actr.MergeParams(params)
	return s.actorRepo.Update(ctx, actr)
}

func (s *service) DeleteActor(ctx context.Context, id int) error {
	return s.actorRepo.Delete(ctx, id)
}

func (s *service) OrderActors(ctx context.Context, limit, offset int) ([]*entity.ActorWithMovies, error) {
	actors, err := s.actorRepo.OrderActors(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	actorIDs := lo.Map(actors, func(actor *entity.Actor, _ int) int {
		return actor.ID
	})

	movies, err := s.moviesRepo.GetMoviesByActors(ctx, actorIDs)
	if err != nil {
		return nil, err
	}

	actorWithMoview := make([]*entity.ActorWithMovies, 0, len(actors))
	for _, actor := range actors {
		actorWithMoview = append(actorWithMoview, &entity.ActorWithMovies{
			Actor:  actor,
			Movies: movies[actor.ID],
		})
	}

	return actorWithMoview, nil
}

func (s *service) GetActor(ctx context.Context, id int) (*entity.ActorWithMovies, error) {
	actor, err := s.actorRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	movies, err := s.moviesRepo.GetMoviesByActorID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.ActorWithMovies{
		Actor:  actor,
		Movies: movies,
	}, nil
}
