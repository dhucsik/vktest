package movie

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/repository"
	"github.com/samber/lo"
)

type Service interface {
	CreateMovie(ctx context.Context, movie *entity.Movie, actorIDs []int) (int, error)
	UpdateMovie(ctx context.Context, id int, params *entity.UpdateMovieParams) error
	DeleteMovie(ctx context.Context, id int) error
	GetMovie(ctx context.Context, id int) (*entity.MovieWithActors, error)
	SearchByTitle(ctx context.Context, search string, limit, offset int) ([]*entity.MovieWithActors, error)
	SearchByActor(ctx context.Context, actor string, limit, offset int) ([]*entity.MovieWithActors, error)
	OrderMovies(ctx context.Context, order string, limit, offset int) ([]*entity.MovieWithActors, error)
}

type service struct {
	movieRepo repository.MovieRepository
	actorRepo repository.ActorRepository
}

func NewService(
	movieRepo repository.MovieRepository,
	actorRepo repository.ActorRepository,
) Service {
	return &service{
		movieRepo: movieRepo,
		actorRepo: actorRepo,
	}
}

func (s *service) CreateMovie(ctx context.Context, movie *entity.Movie, actorIDs []int) (int, error) {
	return s.movieRepo.Create(ctx, movie, actorIDs)
}

func (s *service) UpdateMovie(ctx context.Context, id int, params *entity.UpdateMovieParams) error {
	movie, err := s.movieRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	actorIDs := movie.MergeParams(params)

	return s.movieRepo.Update(ctx, movie, actorIDs)
}

func (s *service) DeleteMovie(ctx context.Context, id int) error {
	return s.movieRepo.Delete(ctx, id)
}

func (s *service) GetMovie(ctx context.Context, id int) (*entity.MovieWithActors, error) {
	movie, err := s.movieRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	actors, err := s.actorRepo.GetActorsByMovie(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.MovieWithActors{
		Movie:  movie,
		Actors: actors,
	}, nil
}

func (s *service) SearchByTitle(ctx context.Context, search string, limit, offset int) ([]*entity.MovieWithActors, error) {
	movies, err := s.movieRepo.Search(ctx, search, limit, offset)
	if err != nil {
		return nil, err
	}

	movieIDs := lo.Map(movies, func(movie *entity.Movie, _ int) int {
		return movie.ID
	})

	actors, err := s.actorRepo.GetActorsByMovies(ctx, movieIDs)
	if err != nil {
		return nil, err
	}

	movieWithActors := make([]*entity.MovieWithActors, 0, len(movies))
	for _, movie := range movies {
		movieWithActors = append(movieWithActors, &entity.MovieWithActors{
			Movie:  movie,
			Actors: actors[movie.ID],
		})
	}

	return movieWithActors, nil
}

func (s *service) SearchByActor(ctx context.Context, actor string, limit, offset int) ([]*entity.MovieWithActors, error) {
	actors, err := s.actorRepo.SearchActors(ctx, actor, limit, offset)
	if err != nil {
		return nil, err
	}

	actorIDs := lo.Map(actors, func(actor *entity.Actor, _ int) int {
		return actor.ID
	})

	movies, err := s.movieRepo.GetMoviesByActorIDs(ctx, actorIDs)
	if err != nil {
		return nil, err
	}

	movieIDs := lo.Map(movies, func(movie *entity.Movie, _ int) int {
		return movie.ID
	})

	actorsByMovie, err := s.actorRepo.GetActorsByMovies(ctx, movieIDs)
	if err != nil {
		return nil, err
	}

	movieWithActors := make([]*entity.MovieWithActors, 0, len(movies))
	for _, movie := range movies {
		movieWithActors = append(movieWithActors, &entity.MovieWithActors{
			Movie:  movie,
			Actors: actorsByMovie[movie.ID],
		})
	}

	return movieWithActors, nil
}

func (s *service) OrderMovies(ctx context.Context, order string, limit, offset int) ([]*entity.MovieWithActors, error) {
	var movies []*entity.Movie
	var err error

	switch order {
	case "rating":
		movies, err = s.movieRepo.OrderByRating(ctx, limit, offset)
	case "title":
		movies, err = s.movieRepo.OrderByTitle(ctx, limit, offset)
	case "release_date":
		movies, err = s.movieRepo.OrderByReleaseDate(ctx, limit, offset)
	default:
		movies, err = s.movieRepo.OrderByRating(ctx, limit, offset)
	}

	if err != nil {
		return nil, err
	}

	movieIDs := lo.Map(movies, func(movie *entity.Movie, _ int) int {
		return movie.ID
	})

	actors, err := s.actorRepo.GetActorsByMovies(ctx, movieIDs)
	if err != nil {
		return nil, err
	}

	movieWithActors := make([]*entity.MovieWithActors, 0, len(movies))
	for _, movie := range movies {
		movieWithActors = append(movieWithActors, &entity.MovieWithActors{
			Movie:  movie,
			Actors: actors[movie.ID],
		})
	}

	return movieWithActors, nil
}
