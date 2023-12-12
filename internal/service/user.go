package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/repository"
	"github.com/yizeng/toggl-test-signer/pkg/hashing"
)

type UserService struct {
	testRepo *repository.TestRepository
}

func NewUserService(testRepo *repository.TestRepository) *UserService {
	return &UserService{
		testRepo: testRepo,
	}
}

func (s *UserService) SignTest(ctx context.Context, userID uint64, answers []domain.Answer) (string, error) {
	now := time.Now().Unix()
	signature := hashing.CreateSHA256(fmt.Sprintf("%v:%v", userID, now))

	answersJSON, err := json.Marshal(answers)
	if err != nil {
		return "", fmt.Errorf("serialize answers json.Marshal(answers) -> %w", err)
	}

	err = s.testRepo.Create(ctx, userID, answersJSON, signature, now)
	if err != nil {
		return "", fmt.Errorf("create test s.testRepo.Create() -> %w", err)
	}

	return signature, nil
}
