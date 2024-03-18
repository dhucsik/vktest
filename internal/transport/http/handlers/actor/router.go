package actor

import (
	"net/http"

	"github.com/dhucsik/vktest/internal/service/actor"
	"github.com/dhucsik/vktest/internal/transport/http/middlewares"
)

type Controller struct {
	auth   *middlewares.AuthMiddleware
	logger *middlewares.LoggerMiddleware

	actorService actor.Service
}

func NewController(
	auth *middlewares.AuthMiddleware,
	logger *middlewares.LoggerMiddleware,
	actorService actor.Service,
) *Controller {
	return &Controller{
		auth:         auth,
		actorService: actorService,
	}
}

func (c *Controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/actors",
		c.logger.Handler(c.auth.Handler(c.createActorHandler)))
	mux.HandleFunc("PUT /api/v1/actors/{id}",
		c.logger.Handler(c.auth.Handler(c.updateActorHandler)))
	mux.HandleFunc("DELETE /api/v1/actors/{id}",
		c.logger.Handler(c.auth.Handler(c.deleteActorHandler)))

	mux.HandleFunc("GET /api/v1/actors",
		c.logger.Handler(c.auth.Handler(c.listActorsHandler)))
	mux.HandleFunc("GET /api/v1/actors/{id}",
		c.logger.Handler(c.auth.Handler(c.getActorByIDHandler)))
}
