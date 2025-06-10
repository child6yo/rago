package main

import (
	"github.com/child6yo/rago/services/generator/internal/app"
	"github.com/child6yo/rago/services/generator/internal/config"
)

func main() {
	cfg := config.InitConfig()
	app := app.CreateApplication(cfg)
	app.StartApplication()
}
