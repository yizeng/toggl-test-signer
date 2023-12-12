package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

func CreateSHA256(input string) string {
	plainText := []byte(input)
	sha256Hash := sha256.Sum256(plainText)

	return hex.EncodeToString(sha256Hash[:])
}
