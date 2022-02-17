package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	middlewares "github.com/lbergamim-daitan/golang-rump-up/internal/middleware"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
	Auth   bool
}

func Config(r *mux.Router) {
	routes := CompaniesRoute
	routes = append(routes, PhonesRoute...)

	for _, route := range routes {
		if route.Auth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Func))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}
}
