package router

import (
	"github.com/asnell/client-server/requestlog"
	"github.com/asnell/client-server/server"
	"github.com/go-chi/chi"
)

func New(s *server.Server) *chi.Mux {
	l := s.Logger()
	r := chi.NewRouter()

	r.Method("GET", "/", requestlog.NewHandler(server.HandleIndex, l))

	return r
}
