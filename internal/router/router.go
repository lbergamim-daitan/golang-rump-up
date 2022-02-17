package router

import (
	"github.com/gorilla/mux"
	"github.com/lbergamim-daitan/golang-rump-up/internal/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	routes.Config(r)
	return r
}
