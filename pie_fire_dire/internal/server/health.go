package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type livenessStats struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type readinessStats struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func healthz(e *echo.Echo) {
	e.GET("/liveness", liveness)
	e.GET("/readiness", readiness)
}

func liveness(c echo.Context) error {
	liveness := livenessStats{
		Code:    "success",
		Message: "Liveness healthy",
	}

	return c.JSON(http.StatusOK, liveness)
}

func readiness(c echo.Context) error {
	readiness := livenessStats{
		Code:    "success",
		Message: "Readiness healthy",
	}

	return c.JSON(http.StatusOK, readiness)
}
