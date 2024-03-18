package user

import (
	"github.com/dhucsik/vktest/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type userModel struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
}

func convertUser(u *entity.User) (*userModel, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &userModel{
		ID:       u.ID,
		Username: u.Username,
		Password: string(hash),
		Role:     string(u.Role),
	}, nil
}

func (u *userModel) convert() *entity.User {
	return &entity.User{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Role:     u.Role,
	}
}
