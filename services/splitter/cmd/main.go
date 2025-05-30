package main

import (
	"github.com/child6yo/rago/services/splitter/internal/app"
	"github.com/child6yo/rago/services/splitter/internal/config"
)

func main() {
	cfg := config.InitConfig()
	app := app.CreateApplication(cfg)
	app.StartApplication()
}