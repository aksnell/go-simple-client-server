package main

import (
	"fmt"
	"github.com/asnell/client-server/config"
	"github.com/asnell/client-server/router"
	"github.com/asnell/client-server/server"
	"github.com/asnell/client-server/util/logger"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {

}

func main() {
	appConfig := config.AppConfig()

	l := logger.New(appConfig.Debug)

	app := server.New(l)

	r := router.New(app)

	address := fmt.Sprintf(":%d", appConfig.Server.Port)

	l.Info().Msgf("Starting server :%v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      r,
		ReadTimeout:  appConfig.Server.TimeoutRead,
		WriteTimeout: appConfig.Server.TimeoutWrite,
		IdleTimeout:  appConfig.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		l.Fatal().Err(err).Msg("Server startup failed!")
	}
}
