package postgresql

import (
	"time"

	"github.com/lib/pq"
)

const (
	sqlEventTable = "events"
)

type sqlEvent struct {
	ID       string         `db:"id"`
	Name     string         `db:"name"`
	Date     time.Time      `db:"date"`
	Shown    int            `db:"shown"`
	Keywords pq.StringArray `db:"keywords"`
}
