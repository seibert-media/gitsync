package gitsync

import (
	"fmt"
	"net/http"

	"github.com/seibert-media/gitsync/pkg/handler"

	"github.com/siddontang/go/log"
	"go.uber.org/zap"
)

// Server creates all required components and starts the http server
type Server struct {
	Port int
}

// PrepareAndServe the handler
func (s *Server) PrepareAndServe() error {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: &handler.Syncer{
			Git:  nil,
			Hook: nil,
		},
	}
	log.Info("listening", zap.Int("port", s.Port))
	return server.ListenAndServe()
}
