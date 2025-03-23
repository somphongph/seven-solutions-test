package handlers

import "pie-fire-dire/internal/api/beef/services"

type handlers struct {
	service services.Servicer
}

func NewHandlers(service services.Servicer) *handlers {
	return &handlers{
		service: service,
	}
}
