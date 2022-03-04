package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lbergamim-daitan/golang-rump-up/internal/config"
	"github.com/lbergamim-daitan/golang-rump-up/internal/router"
)

func main() {
	config.Load()
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
