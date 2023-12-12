package v1

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/request"
	"github.com/yizeng/toggl-test-signer/internal/web/handler/v1/response"
	"go.uber.org/zap"
)

type UserService interface {
	SignTest(ctx context.Context, userID uint64, answers []domain.Answer) (string, error)
}

type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) HandleSignTest(w http.ResponseWriter, r *http.Request) {
	userID, err := h.getUserIDFromJWT(r.Context())
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest(err))

		return
	}

	req := request.SignTest{}
	if err := render.Bind(r, &req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, ErrBadRequest(err))

		return
	}

	signature, err := h.svc.SignTest(r.Context(), userID, req.Answers)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, ErrInternalServerError(err))

		return
	}

	resp := response.SignTest{
		UserID:    userID,
		Signature: signature,
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, resp)
}

func (h *UserHandler) getUserIDFromJWT(ctx context.Context) (uint64, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		err := errors.New("retrieving JWT from context failed")
		zap.L().Error(err.Error(), zap.Error(err))

		return 0, err
	}

	claimUserID, ok := claims["userID"]
	if !ok {
		err := errors.New("userID is not found in JWT payload")
		zap.L().Error(err.Error(), zap.Error(err))

		return 0, err
	}

	userID, ok := claimUserID.(float64)
	if !ok {
		err := errors.New("userID cannot be casted into int")
		zap.L().Error(err.Error(), zap.Error(err))

		return 0, err
	}

	return uint64(userID), nil
}
