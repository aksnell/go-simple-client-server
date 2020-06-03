package server

import (
	"fmt"
	"github.com/asnell/client-server/util/logger"
	"net/http"
)

type Server struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *Server {
	return &Server{logger: logger}
}

func (s *Server) Logger() *logger.Logger {
	return s.logger
}

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Write([]byte("Hello World!"))
}
