package api

import (
	"github.com/go-chi/chi/v5"

	"github.com/ljgago/sift/internal/repository"

	"github.com/ljgago/sift/api/packages"
	"github.com/ljgago/sift/api/search"
	"github.com/ljgago/sift/api/stars"
)

func RegistryRoutes(r chi.Router, repo *repository.Repository) {
	r.Route("/api", func(r chi.Router) {
		search.RegisterRoutes(r, repo.Registry)
		packages.RegisterRoutes(r, repo.Registry)
		stars.RegisterRoutes(r, repo.Github)
	})
}
