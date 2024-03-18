package movie

import (
	"net/http"

	"github.com/dhucsik/vktest/internal/service/movie"
	"github.com/dhucsik/vktest/internal/transport/http/middlewares"
)

type Controller struct {
	auth   *middlewares.AuthMiddleware
	logger *middlewares.LoggerMiddleware

	moviesService movie.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	logger *middlewares.LoggerMiddleware,
	moviesService movie.Service,
) *Controller {
	return &Controller{
		auth:          auth,
		logger:        logger,
		moviesService: moviesService,
	}
}

func (c *Controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/movies",
		c.logger.Handler(c.auth.Handler(c.createMovieHandler)))
	mux.HandleFunc("PUT /api/v1/movies/{id}",
		c.logger.Handler(c.auth.Handler(c.updateMovieHandler)))
	mux.HandleFunc("DELETE /api/v1/movies/{id}",
		c.logger.Handler(c.auth.Handler(c.deleteMovieHandler)))

	mux.HandleFunc("GET /api/v1/movies/search",
		c.logger.Handler(c.auth.Handler(c.searchMoviesHandler)))
	mux.HandleFunc("GET /api/v1/movies/{id}",
		c.logger.Handler(c.auth.Handler(c.getMovieByIDHandler)))
	mux.HandleFunc("GET /api/v1/movies",
		c.logger.Handler(c.auth.Handler(c.listMoviesHandler)))
}
