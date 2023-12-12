package service

import (
	"context"

	"github.com/yizeng/toggl-test-signer/internal/domain"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SignAnswers(ctx context.Context, userID uint64, answers []domain.Answer) (string, error) {
	// TODO create userID + answers
	// Generate a signature and save
	// Return signature
	return "", nil
}
