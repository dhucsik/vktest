package app

import (
	"context"

	"github.com/dhucsik/vktest/config"
	"github.com/dhucsik/vktest/internal/envs"
	"github.com/dhucsik/vktest/internal/repository"
	"github.com/dhucsik/vktest/internal/service/actor"
	"github.com/dhucsik/vktest/internal/service/movie"
	"github.com/dhucsik/vktest/internal/service/user"
	"github.com/dhucsik/vktest/internal/transport/http"
)

type Application struct {
	cfg    *config.Config
	server *http.Server

	Repository *repository.Repository

	actorService actor.Service
	movieService movie.Service
	userService  user.Service
}

func NewApplication(ctx context.Context) (*Application, error) {
	cfg, err := config.Parse()
	if err != nil {
		return nil, err
	}

	return &Application{
		cfg: cfg,
	}, nil
}

func InitApp(ctx context.Context) (*Application, error) {
	app, err := NewApplication(ctx)
	if err != nil {
		return nil, err
	}

	for _, init := range []func(context.Context) error{
		app.initRepository,
		app.initServices,
		app.initServer,
	} {
		if err := init(ctx); err != nil {
			return nil, err
		}
	}

	return app, nil
}

func (a *Application) initRepository(ctx context.Context) error {
	var err error
	a.Repository, err = repository.NewRepository(ctx, a.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (a *Application) initServices(ctx context.Context) error {
	var err error

	a.actorService = actor.NewService(a.Repository.Actor, a.Repository.Movie)
	a.movieService = movie.NewService(a.Repository.Movie, a.Repository.Actor)
	a.userService, err = user.NewService(
		a.Repository.User,
		a.cfg.Env.Get(envs.AccessTTL),
		a.cfg.Env.Get(envs.RefreshTTL),
	)

	if err != nil {
		return err
	}

	return nil
}

func (a *Application) initServer(ctx context.Context) error {
	a.server = http.NewServer(
		a.cfg.HTTP,
		a.actorService,
		a.movieService,
		a.userService,
	)

	return nil
}

func (a *Application) Start(ctx context.Context) error {
	return a.server.Start()
}
