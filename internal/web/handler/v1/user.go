package v1

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/request"
)

type UserService interface {
	SignAnswers(ctx context.Context, userID uint64, answers []domain.Answer) (string, error)
}

type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) HandleSignAnswers(w http.ResponseWriter, r *http.Request) {
	var userID uint64 = 0 // TODO: Read from JWT.

	req := request.SignAnswers{}
	if err := render.Bind(r, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest(err))

		return
	}

	answers, err := h.svc.SignAnswers(r.Context(), userID, req.Answers)
	if err != nil {
		// TODO: Handle explicitly if user not found.
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrInternalServerError(err))
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, answers)
}
