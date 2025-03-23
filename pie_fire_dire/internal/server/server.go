package server

import (
	"fmt"
	"net/http"
	"pie-fire-dire/configs"

	"github.com/labstack/gommon/log"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	conf := configs.GetConfig()

	newServer := &Server{
		port: conf.App.Port,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      newServer.RegisterRoutes(),
		IdleTimeout:  conf.Server.IdleTimeout,
		ReadTimeout:  conf.Server.ReadTimeout,
		WriteTimeout: conf.Server.WriteTimeout,
	}

	log.Infof("Server is running on port: %d", conf.App.Port)

	return server
}
