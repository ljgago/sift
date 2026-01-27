package search

import (
	"net/http"

	"github.com/ljgago/sift/internal/log"
	"github.com/ljgago/sift/internal/rest"
)

type Handler struct {
	service *Service
}

func (h *Handler) ListPackages(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.RawQuery

	response, err := h.service.Packages(ctx, query)
	if err != nil {
		log.Errorf("search Handler ListPackage: %s", err.Error())
		rest.ErrorResponse(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	rest.SuccessResponse(w, http.StatusOK, response)
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
