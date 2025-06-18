package server

import (
	"context"
	"net/http"
)

// Server определяет структуру HTTP сервера.
type Server struct {
	httpServer *http.Server
}

// Run запускает HTTP сервер.
func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    0,
		WriteTimeout:   0,
	}

	return s.httpServer.ListenAndServe()
}

// Shutdown останавливает HTTP сервер.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
