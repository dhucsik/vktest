package actor

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lib/pq"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, actor *entity.Actor) (int, error) {
	model := convertActor(actor)

	row := r.db.QueryRow(ctx, createActorStmt, model.FullName, model.Gender, model.DateOfBirth)
	err := row.Scan(&model.ID)
	if err != nil {
		if err != nil {
			return 0, err
		}
	}

	return model.ID, nil
}

func (r *Repository) Update(ctx context.Context, actor *entity.Actor) error {
	model := convertActor(actor)

	_, err := r.db.Exec(ctx, updateActorStmt, model.ID, model.FullName, model.Gender, model.DateOfBirth)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*entity.Actor, error) {
	row := r.db.QueryRow(ctx, getActorByIDStmt, id)

	var model actorModel
	err := row.Scan(&model.ID, &model.FullName, &model.Gender, &model.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return model.convert(), nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, deleteActorStmt, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetActorsByMovie(ctx context.Context, movieID int) ([]*entity.Actor, error) {
	rows, err := r.db.Query(ctx, getActorsByMovieStmt, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []*entity.Actor
	for rows.Next() {
		var model actorModel
		err := rows.Scan(&model.ID, &model.FullName, &model.Gender, &model.DateOfBirth)
		if err != nil {
			return nil, err
		}

		actors = append(actors, model.convert())
	}

	return actors, nil
}

func (r *Repository) GetActorsByMovies(ctx context.Context, movieIDs []int) (map[int][]*entity.Actor, error) {
	rows, err := r.db.Query(ctx, getActorsByMoviesStmt, pq.Array(movieIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	actorsByMovie := make(map[int][]*entity.Actor)
	for rows.Next() {
		var movieID int
		var model actorModel
		err := rows.Scan(&model.ID, &model.FullName, &model.Gender, &model.DateOfBirth, &movieID)
		if err != nil {
			return nil, err
		}

		actors, ok := actorsByMovie[movieID]
		if !ok {
			actors = make([]*entity.Actor, 0)
		}

		actorsByMovie[movieID] = append(actors, model.convert())
	}

	return actorsByMovie, nil
}

func (r *Repository) SearchActors(ctx context.Context, actor string, limit, offset int) ([]*entity.Actor, error) {
	rows, err := r.db.Query(ctx, searchActorsStmt, actor, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []*entity.Actor
	for rows.Next() {
		var model actorModel
		err := rows.Scan(&model.ID, &model.FullName, &model.Gender, &model.DateOfBirth)
		if err != nil {
			return nil, err
		}

		actors = append(actors, model.convert())
	}

	return actors, nil
}

func (r *Repository) OrderActors(ctx context.Context, limit, offset int) ([]*entity.Actor, error) {
	rows, err := r.db.Query(ctx, orderActorsStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []*entity.Actor
	for rows.Next() {
		var model actorModel
		err := rows.Scan(&model.ID, &model.FullName, &model.Gender, &model.DateOfBirth)
		if err != nil {
			return nil, err
		}

		actors = append(actors, model.convert())
	}

	return actors, nil
}
