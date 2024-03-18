package movie

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/lib/pq"
)

func (r *Repository) GetByID(ctx context.Context, id int) (*entity.Movie, error) {
	row := r.db.QueryRow(ctx, getByIDStmt, id)

	var model movieModel
	err := row.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
	if err != nil {
		return nil, err
	}

	return model.convert(), nil
}

func (r *Repository) Search(ctx context.Context, search string, limit, offset int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, searchMoviesStmt, search, limit, offset)
	if err != nil {
		return nil, err
	}

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel
		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}

func (r *Repository) GetMoviesByActorIDs(ctx context.Context, actorIDs []int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, getMoviesByActorIDsStmt, pq.Array(actorIDs))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel

		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}

func (r *Repository) GetMoviesByActors(ctx context.Context, actorID []int) (map[int][]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, getMoviesByActorsStmt, pq.Array(actorID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	moviesByActor := make(map[int][]*entity.Movie)
	for rows.Next() {
		var movieID int
		var model movieModel
		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating, &movieID)
		if err != nil {
			return nil, err
		}

		movies, ok := moviesByActor[movieID]
		if !ok {
			movies = make([]*entity.Movie, 0)
		}

		moviesByActor[movieID] = append(movies, model.convert())
	}

	return moviesByActor, nil
}

func (r *Repository) OrderByRating(ctx context.Context, limit, offset int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, orderMoviesByRatingStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel

		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}

func (r *Repository) OrderByTitle(ctx context.Context, limit, offset int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, orderMoviesByTitleStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel

		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}

func (r *Repository) OrderByReleaseDate(ctx context.Context, limit, offset int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, orderMoviesByReleaseDateStmt, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel

		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}

func (r *Repository) GetMoviesByActorID(ctx context.Context, actorID int) ([]*entity.Movie, error) {
	rows, err := r.db.Query(ctx, getMoviesByActorID, actorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*entity.Movie
	for rows.Next() {
		var model movieModel

		err = rows.Scan(&model.ID, &model.Title, &model.Description, &model.ReleaseDate, &model.Rating)
		if err != nil {
			return nil, err
		}

		movies = append(movies, model.convert())
	}

	return movies, nil
}
