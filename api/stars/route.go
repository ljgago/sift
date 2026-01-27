package stars

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, r Repository) {
	s := NewService(r)
	h := NewHandler(s)

	router.Get("/github/repos/{owner}/{repo}/stars", h.GetStars)
}
