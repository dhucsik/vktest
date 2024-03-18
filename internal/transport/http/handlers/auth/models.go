package auth

import (
	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/enum"
)

type registerRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *registerRequest) convert() (*entity.User, error) {
	return &entity.User{
		Username: r.Username,
		Password: r.Password,
		Role:     enum.UserRoleUser,
	}, nil
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type refreshRequest struct {
	Token string `json:"token"`
}

type errorResponse struct {
	Error string `json:"error"`
}
