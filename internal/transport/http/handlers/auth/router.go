package auth

import (
	"net/http"

	"github.com/dhucsik/vktest/internal/service/user"
	"github.com/dhucsik/vktest/internal/transport/http/middlewares"
)

type Controller struct {
	logger       *middlewares.LoggerMiddleware
	usersService user.Service
}

func NewController(
	usersService user.Service,
	logger *middlewares.LoggerMiddleware,
) *Controller {
	return &Controller{
		usersService: usersService,
		logger:       logger,
	}
}

func (c *Controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/auth/register",
		c.logger.Handler(c.signupHandler))
	mux.HandleFunc("POST /api/v1/auth/login",
		c.logger.Handler(c.loginHandler))
	mux.HandleFunc("POST /api/v1/auth/refresh",
		c.logger.Handler(c.refreshHandler))
}
