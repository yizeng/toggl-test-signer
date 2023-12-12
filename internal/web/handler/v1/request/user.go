package request

import (
	"net/http"

	"github.com/yizeng/toggl-test-signer/internal/domain"
)

type SignAnswers struct {
	Answers []domain.Answer
}

func (c *SignAnswers) Bind(r *http.Request) error {
	return nil
}
