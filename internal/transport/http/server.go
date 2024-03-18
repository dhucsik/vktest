package http

import (
	"fmt"
	"net/http"

	"github.com/dhucsik/vktest/config"
	"github.com/dhucsik/vktest/internal/service/actor"
	"github.com/dhucsik/vktest/internal/service/movie"
	"github.com/dhucsik/vktest/internal/service/user"
	actorC "github.com/dhucsik/vktest/internal/transport/http/handlers/actor"
	"github.com/dhucsik/vktest/internal/transport/http/handlers/auth"
	movieC "github.com/dhucsik/vktest/internal/transport/http/handlers/movie"
	"github.com/dhucsik/vktest/internal/transport/http/handlers/swag"
	"github.com/dhucsik/vktest/internal/transport/http/middlewares"
)

// @title Swagger Movies API
// @version 1.0
// @description movies server.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.htm

// @host localhost:8080
// @BasePath /api/v1

type Server struct {
	actorsService actor.Service
	moviesService movie.Service
	userService   user.Service

	srv  *http.ServeMux
	addr string
}

type IController interface {
	Init(mux *http.ServeMux)
}

func NewServer(
	cfg *config.HTTP,
	actorsService actor.Service,
	moviesService movie.Service,
	userService user.Service,
) *Server {
	srv := http.NewServeMux()

	server := &Server{
		srv:           srv,
		actorsService: actorsService,
		moviesService: moviesService,
		userService:   userService,
		addr:          fmt.Sprintf(":%s", cfg.Port),
	}
	server.init()

	return server
}

func (s *Server) init() {
	authMiddleware := middlewares.NewAuthMiddleware()
	loggerMiddleware := middlewares.NewLoggerMiddleware()

	s.WithControllers(
		actorC.NewController(authMiddleware, loggerMiddleware, s.actorsService),
		movieC.NewController(authMiddleware, loggerMiddleware, s.moviesService),
		auth.NewController(s.userService, loggerMiddleware),
		swag.NewController(),
	)
}

func (s *Server) WithControllers(controllers ...IController) {
	for _, c := range controllers {
		c.Init(s.srv)
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.srv)
}
