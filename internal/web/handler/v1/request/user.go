package request

import (
	"net/http"

	"github.com/yizeng/toggl-test-signer/internal/domain"
)

type SignTest struct {
	Answers []domain.Answer
}

func (c *SignTest) Bind(r *http.Request) error {
	return nil
}
