package routes

import (
	"net/http"

	"github.com/lbergamim-daitan/golang-rump-up/internal/controllers"
)

var PhonesRoute = []Route{
	{
		URI:    "/phones",
		Method: http.MethodGet,
		Func:   controllers.ListPhones,
		Auth:   true,
	},
}
