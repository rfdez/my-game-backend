package postgresql

const (
	sqlEventQuestionTable = "event_questions"
)

type sqlEventQuestion struct {
	EventID    string `db:"event_id"`
	QuestionID string `db:"question_id"`
	Round      int    `db:"round"`
}
