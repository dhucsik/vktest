package swag

import (
	"net/http"

	_ "github.com/dhucsik/vktest/swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Init(mux *http.ServeMux) {
	mux.HandleFunc("GET /swagger/*", httpSwagger.Handler())
}
