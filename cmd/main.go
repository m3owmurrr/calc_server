package main

import (
	"github.com/m3owmurrr/calc/internal/application"
	"github.com/m3owmurrr/calc/internal/config"
)

func main() {
	conf := config.GetConfig()
	app := application.NewApplication(conf)
	app.Run()
}
