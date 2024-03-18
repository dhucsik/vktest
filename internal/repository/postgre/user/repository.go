package user

import (
	"context"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(ctx context.Context, user *entity.User) error {
	model, err := convertUser(user)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, createUserStmt, model.Username, model.Password, model.Role)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	row := r.db.QueryRow(ctx, getUserByIDStmt, id)

	var model userModel
	err := row.Scan(&model.ID, &model.Username, &model.Password, &model.Role)
	if err != nil {
		return nil, err
	}

	return model.convert(), nil
}

func (r *Repository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	row := r.db.QueryRow(ctx, getUserByUsernameStmt, username)

	var model userModel
	err := row.Scan(&model.ID, &model.Username, &model.Password, &model.Role)
	if err != nil {
		return nil, err
	}

	return model.convert(), nil
}
