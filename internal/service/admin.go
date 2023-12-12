package service

import (
	"context"

	"github.com/yizeng/toggl-test-signer/internal/domain"
)

type AdminService struct{}

func NewAdminService() *AdminService {
	return &AdminService{}
}

func (s *AdminService) VerifySignature(ctx context.Context, userID uint64, signature string) ([]domain.Answer, error) {
	// TODO find user by userID in DB.
	// Compare th signature
	// Return answers
	return nil, nil
}
