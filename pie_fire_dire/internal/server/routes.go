package server

import (
	beef "pie-fire-dire/internal/api/beef/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() *echo.Echo {
	e := echo.New()

	secureConfig := middleware.SecureConfig{
		ContentTypeNosniff: "nosniff",
		ReferrerPolicy:     "same-origin",
	}

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.SecureWithConfig(secureConfig),
	)

	// health
	healthz(e)

	apiRoot := e.Group("/v1")

	// places
	beef.NewRoutes(apiRoot)

	return e
}
