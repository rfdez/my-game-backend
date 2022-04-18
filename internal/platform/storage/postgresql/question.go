package postgresql

const (
	sqlQuestionTable = "questions"
)

type sqlQuestion struct {
	ID   string `db:"id"`
	Text string `db:"text"`
}
