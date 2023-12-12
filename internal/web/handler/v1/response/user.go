package response

type SignTest struct {
	UserID    uint64 `json:"userID"`
	Signature string `json:"signature"`
}
