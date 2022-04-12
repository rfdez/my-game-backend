package postgresql

const (
	sqlEventTable = "events"
)

type sqlEvent struct {
	ID   string `db:"id"`
	Name string `db:"name"`
}
