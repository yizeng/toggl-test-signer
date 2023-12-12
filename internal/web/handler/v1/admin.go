package v1

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/service"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/request"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/response"
)

type AdminService interface {
	VerifySignature(ctx context.Context, userID uint64, signature string) ([]domain.Answer, int64, error)
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

	answers, signedAt, err := h.svc.VerifySignature(r.Context(), req.UserID, req.Signature)
	if err != nil {
		if errors.Is(err, service.ErrTestNotFound) {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, ErrNotFound(err))

			return
		} else {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, ErrInternalServerError(err))

			return
		}
	}

	resp := response.VerifySignature{
		Answers:  answers,
		SignedAt: time.Unix(signedAt, 0),
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}
