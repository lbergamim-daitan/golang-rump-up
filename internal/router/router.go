package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lbergamim-daitan/golang-rump-up/internal/router/routes"
)

func Generate() http.Handler {
	r := mux.NewRouter()
	handler := routes.Config(r)
	return handler
}
