package movie

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
)

func (r *Repository) Create(ctx context.Context, movie *entity.Movie, actorIDs []int) (int, error) {
	model := convertMovie(movie)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	row := tx.QueryRow(ctx, createMovieStmt, model.Title, model.Description, model.ReleaseDate, model.Rating)
	err = row.Scan(&model.ID)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	for _, actorID := range actorIDs {
		_, err = tx.Exec(ctx, createMovieActorStmt, model.ID, actorID)
		if err != nil {
			tx.Rollback(ctx)
			return 0, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return 0, err
	}

	return model.ID, nil
}

func (r *Repository) Update(ctx context.Context, movie *entity.Movie, actorIDs []int) error {
	model := convertMovie(movie)

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, updateMovieStmt, model.ID, model.Title, model.Description, model.ReleaseDate, model.Rating)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, deleteMovieActorsStmt, model.ID)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	for _, actorID := range actorIDs {
		_, err = tx.Exec(ctx, createMovieActorStmt, model.ID, actorID)
		if err != nil {
			tx.Rollback(ctx)
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, deleteMovieActorsStmt, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	_, err = tx.Exec(ctx, deleteMovieStmt, id)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
