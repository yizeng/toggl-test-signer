package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yizeng/toggl-test-signer/internal/domain"
	"github.com/yizeng/toggl-test-signer/internal/repository/dao"
)

var (
	ErrTestNotFound = dao.ErrTestNotFound
)

type TestRepository struct {
	dao dao.TestDAO
}

func NewTestRepository(dao dao.TestDAO) *TestRepository {
	return &TestRepository{
		dao: dao,
	}
}

func (r TestRepository) Create(ctx context.Context, userID uint64, answers []byte, signature string, signedAt int64) error {
	err := r.dao.Create(ctx, userID, answers, signature, signedAt)
	if err != nil {
		return fmt.Errorf("create test r.dao.Create() -> %w", err)
	}

	return nil
}

func (r TestRepository) Find(ctx context.Context, userID uint64, signature string) ([]domain.Answer, error) {
	answers, err := r.dao.Find(ctx, userID, signature)
	if err != nil {
		return nil, fmt.Errorf("find test r.dao.Find() -> %w", err)
	}

	var result []domain.Answer
	err = json.Unmarshal(answers, &result)
	if err != nil {
		return nil, fmt.Errorf(" unmarshal answers json.Unmarshal -> %w", err)
	}

	return result, nil
}
