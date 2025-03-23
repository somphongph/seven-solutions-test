package routes

import (
	"pie-fire-dire/internal/api/beef/handlers"
	"pie-fire-dire/internal/api/beef/services"

	"github.com/labstack/echo/v4"
)

func NewRoutes(apiRoot *echo.Group) {
	service := services.NewServices()

	// beef
	h := handlers.NewHandlers(service)
	g := apiRoot.Group("/beef")
	g.POST("/summary", h.PostSummary)
}
