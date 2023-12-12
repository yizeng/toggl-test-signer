package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/service"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/request"
)

type AdminService interface {
	VerifySignature(ctx context.Context, userID uint64, signature string) ([]domain.Answer, error)
}

type AdminHandler struct {
	svc AdminService
}

func NewAdminHandler(svc *service.AdminService) *AdminHandler {
	return &AdminHandler{
		svc: svc,
	}
}

func (h *AdminHandler) HandleVerifySignature(w http.ResponseWriter, r *http.Request) {
	req := request.VerifySignature{}
	if err := render.Bind(r, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest(err))

		return
	}

	answers, err := h.svc.VerifySignature(r.Context(), req.UserID, req.Signature)
	if err != nil {
		if errors.Is(err, service.ErrTestNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, ErrNotFound(err))
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, ErrInternalServerError(err))
		}
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, answers)
}
