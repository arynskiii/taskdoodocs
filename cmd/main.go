package main

import (
	"doodocs_task/internal/app"
	"doodocs_task/internal/config"

	"github.com/sirupsen/logrus"
)

// 1
func main() {
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatal(err)
	}
	if err := app.App(cfg); err != nil {
		logrus.Fatal(err)
	}
}
