package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
)

type eventQuestionRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewEventQuestionRepository initializes a PostgreSQL implementation of mygame.QuestionRepository.
func NewEventQuestionRepository(db *sql.DB, dbTimeout time.Duration) *eventQuestionRepository {
	return &eventQuestionRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// SearchByEventID implements the mygame.QuestionRepository repository.
func (r *eventQuestionRepository) SearchByEventID(ctx context.Context, eventID mygame.EventID) ([]mygame.EventQuestion, error) {
	eventQuestionSQLStruct := sqlbuilder.NewStruct(new(sqlEventQuestion))

	sb := eventQuestionSQLStruct.SelectFrom(sqlEventQuestionTable)
	sb.Where(sb.E("event_id", eventID.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search questions")
	}
	defer rows.Close()

	var eventQuestions []mygame.EventQuestion
	for rows.Next() {
		var eventQuestion sqlEventQuestion
		if err := rows.Scan(eventQuestionSQLStruct.Addr(&eventQuestion)...); err != nil {
			return nil, errors.Wrap(err, "failed to scan question")
		}

		eq, err := mygame.NewEventQuestion(
			eventQuestion.EventID,
			eventQuestion.QuestionID,
			eventQuestion.Round,
		)
		if err != nil {
			return nil, err
		}

		eventQuestions = append(eventQuestions, eq)
	}

	return eventQuestions, nil
}
