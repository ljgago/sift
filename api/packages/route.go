package packages

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, r Repository) {
	s := NewService(r)
	h := NewHandler(s)

	router.Get("/registry/package/*", h.GetPackage)
}
