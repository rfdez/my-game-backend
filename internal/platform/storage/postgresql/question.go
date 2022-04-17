package postgresql

const (
	sqlQuestionTable = "questions"
)

type sqlQuestion struct {
	ID      string `db:"id"`
	Text    string `db:"text"`
	Round   int    `db:"round"`
	EventID string `db:"event_id"`
}
