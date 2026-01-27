package stars

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ljgago/sift/internal/log"
	"github.com/ljgago/sift/internal/rest"
)

type Handler struct {
	service *Service
}

func (h *Handler) GetStars(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	owner := chi.URLParam(r, "owner")
	repo := chi.URLParam(r, "repo")

	response, err := h.service.Stars(ctx, owner, repo)
	if err != nil {
		log.Errorf("stars Handler GetStars: %s", err.Error())
		rest.ErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	w.Header().Set("Cache-Control", "public, max-age=3600")

	rest.SuccessResponse(w, http.StatusOK, response)
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
