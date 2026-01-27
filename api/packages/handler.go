package packages

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/ljgago/sift/internal/log"
	"github.com/ljgago/sift/internal/rest"
)

type Handler struct {
	service *Service
}

func (h *Handler) GetPackage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := chi.URLParam(r, "*")

	response, err := h.service.Package(ctx, name)
	if err != nil {
		log.Errorf("package Handler GetPackage: %s", err.Error())
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
