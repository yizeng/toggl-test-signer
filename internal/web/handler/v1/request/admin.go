package request

import "net/http"

type VerifySignature struct {
	UserID    uint64 `json:"userID"`
	Signature string `json:"signature"`
}

func (c *VerifySignature) Bind(r *http.Request) error {
	return nil
}
