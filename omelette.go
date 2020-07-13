package main

import (
	"log"
	"net/http"
	"omelette/app/connectors"
	"omelette/app/routes"
	"omelette/config"
)

var cfg *config.Config

func init() {
	cfg, _ = config.LoadConfig()
}

func main() {
	log.Println("\033[32m ⇨ Initializing database hang on..\033[0m")
	connectors.Initialize(cfg)
	log.Println("\033[32m ⇨ Initializing router almost done..\033[0m")
	routes.InitializeRouter()
	log.Println("\033[32m ⇨ http server started at " + cfg.Server.Host + ":" + cfg.Server.Port + "\033[0m")
	http.ListenAndServe(":"+cfg.Server.Port, routes.RouterInstance())
}
