package postgresql

import (
	"context"
	"database/sql"
	"time"
)

type checkRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCheckRepository initializes a PostgreSQL implementation of mygame.CheckRepository.
func NewCheckRepository(db *sql.DB, dbTimeout time.Duration) *checkRepository {
	return &checkRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Status implements mygame.CheckRepository interface.
func (r *checkRepository) Status(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	return r.db.PingContext(ctx)
}
