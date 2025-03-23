package handlers

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handlers) PostSummary(c echo.Context) error {
	body, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return c.String(http.StatusBadRequest, "Failed to read body")
	}

	resp := h.service.CountBeefSummary(string(body))

	return c.JSON(http.StatusOK, resp)
}
