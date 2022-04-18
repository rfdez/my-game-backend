package postgresql

const (
	sqlAnswerTable = "answers"
)

type sqlAnswer struct {
	EventID    string `db:"event_id"`
	QuestionID string `db:"question_id"`
	Text       string `db:"text"`
}
