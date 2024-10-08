// Code generated by raiden-cli; DO NOT EDIT.
package bootstrap

import (
	"github.com/sev-2/raiden"
	"app/internal/controllers"
	"app/internal/models"
	"github.com/valyala/fasthttp"
)

func RegisterRoute(server *raiden.Server) {
	server.RegisterRoute([]*raiden.Route{
		{
			Type:       raiden.RouteTypeCustom,
			Path:       "/hello",
			Methods:    []string{fasthttp.MethodGet},
			Controller: &controllers.HelloWorldController{},
		},
		{
			Type:       raiden.RouteTypeRest,
			Path:       "/tasks",
			Methods:    []string{},
			Controller: &controllers.BooksController{},
			Model:      models.Tasks{},
		},
	})
}
