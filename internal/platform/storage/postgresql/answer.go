package postgresql

const (
	sqlAnswerTable = "answers"
)

type sqlAnswer struct {
	ID         string `db:"id"`
	Text       string `db:"text"`
	QuestionID string `db:"question_id"`
}
