package response

import (
	"time"

	"github.com/yizeng/toggl-test-signer/internal/domain"
)

type VerifySignature struct {
	Answers  []domain.Answer `json:"answers"`
	SignedAt time.Time       `json:"signedAt"`
}
