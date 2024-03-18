package user

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/dhucsik/vktest/internal/entity"
	"github.com/dhucsik/vktest/internal/repository"
	"github.com/dhucsik/vktest/internal/util/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(ctx context.Context, user *entity.User) error
	Auth(ctx context.Context, username, password string) (string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, string, error)
}

type service struct {
	accessTTL  time.Duration
	refreshTTl time.Duration
	userRepo   repository.UserRepository
}

func NewService(
	userRepo repository.UserRepository,
	accessTTLMin string,
	refreshTTLMin string,
) (Service, error) {
	accTTL, err := strconv.Atoi(accessTTLMin)
	if err != nil {
		return nil, err
	}

	refreshTTL, err := strconv.Atoi(refreshTTLMin)
	if err != nil {
		return nil, err
	}

	return &service{
		userRepo:   userRepo,
		accessTTL:  time.Minute * time.Duration(accTTL),
		refreshTTl: time.Minute * time.Duration(refreshTTL),
	}, nil
}

func (s *service) CreateUser(ctx context.Context, user *entity.User) error {
	return s.userRepo.Create(ctx, user)
}

func (s *service) Auth(ctx context.Context, username, password string) (string, string, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", err
	}

	session := &entity.Session{
		UserID: user.ID,
		Role:   user.Role,
	}

	accessToken, err := jwt.GenerateJWT(session, s.accessTTL, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.GenerateJWT(session, s.refreshTTl, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *service) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	session, isRefresh, err := jwt.ParseJWT(refreshToken)
	if err != nil {
		return "", "", err
	}

	if !isRefresh {
		return "", "", errors.New("not a refresh token")
	}

	accessToken, err := jwt.GenerateJWT(session, s.accessTTL, false)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = jwt.GenerateJWT(session, s.refreshTTl, true)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
