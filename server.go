package main

import (
	"log"
	"net/http"

	"github.com/mustanish/omelette/app/config"
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/routes"
)

var cfg *config.Config

func init() {
	cfg, _ = config.LoadConfig()
}

func initConnector() {
	connectors.Initialize(cfg)
}

func initServer() {
	routes.InitializeRouter()
	log.Println("\033[32m" + "⇨ http server started at " + cfg.Server.Host + ":" + cfg.Server.Port + "\033[0m")
	http.ListenAndServe(":"+cfg.Server.Port, routes.RouterInstance())
}

func main() {
	initConnector()
	initServer()
}
