package main

import (
	"github.com/child6yo/rago/api-gateway/internal/app"
	"github.com/child6yo/rago/api-gateway/internal/config"
)

func main() {
	cfg := config.InitConfig()
	a := app.CreateApplication(cfg)

	a.StartApplication()
}