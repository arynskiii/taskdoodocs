package doodocs_task

import (
	"context"
	"doodocs_task/internal/config"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{

		httpServer: &http.Server{
			Addr:              ":" + cfg.Server.Addr,
			Handler:           handler,
			ReadHeaderTimeout: cfg.Server.ReadTimeout,
			WriteTimeout:      cfg.Server.WriteTimeout,
			MaxHeaderBytes:    cfg.Server.MaxHeaderBytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
