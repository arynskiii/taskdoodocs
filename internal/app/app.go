package app

import (
	"context"
	doodocs_task "doodocs_task"
	"doodocs_task/internal/config"
	handler "doodocs_task/internal/delivery"
	"doodocs_task/internal/service"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func App(cfg *config.Config) error {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	service := service.NewService(cfg.SMTPConfig)
	handler := handler.NewHandler(service)
	server := doodocs_task.NewServer(cfg, handler.Init())
	go func() {
		if err := server.Run(); err != nil {
			logrus.Println(err)
		}
	}()
	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("App Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		return err
	}

	return nil
}
