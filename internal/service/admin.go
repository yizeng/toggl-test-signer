package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/repository"
)

var (
	ErrTestNotFound = repository.ErrTestNotFound
)

type AdminService struct {
	testRepo *repository.TestRepository
}

func NewAdminService(testRepo *repository.TestRepository) *AdminService {
	return &AdminService{
		testRepo: testRepo,
	}
}

func (s *AdminService) VerifySignature(ctx context.Context, userID uint64, signature string) ([]domain.Answer, error) {
	answers, err := s.testRepo.Find(ctx, userID, signature)
	if err != nil {
		if errors.Is(err, ErrTestNotFound) {
			return nil, ErrTestNotFound
		}

		return nil, fmt.Errorf("find test s.testRepo.Find() -> %w", err)
	}

	return answers, nil
}
