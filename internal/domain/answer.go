package domain

type Answer struct {
	QuestionID uint64 `json:"questionID"`
	Answer     string `json:"answer"`
}
