package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
)

type questionRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewQuestionRepository initializes a PostgreSQL implementation of mygame.QuestionRepository.
func NewQuestionRepository(db *sql.DB, dbTimeout time.Duration) *questionRepository {
	return &questionRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// SearchByEventID implements the mygame.QuestionRepository repository.
func (r *questionRepository) SearchByEventID(ctx context.Context, eventID mygame.EventID) ([]mygame.Question, error) {
	questionSQLStruct := sqlbuilder.NewStruct(new(sqlQuestion))

	sb := questionSQLStruct.SelectFrom(sqlQuestionTable)
	sb.Where(sb.E("event_id", eventID.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search questions")
	}
	defer rows.Close()

	var questions []mygame.Question
	for rows.Next() {
		var question sqlQuestion
		if err := rows.Scan(questionSQLStruct.Addr(&question)...); err != nil {
			return nil, errors.Wrap(err, "failed to scan question")
		}

		q, err := mygame.NewQuestion(
			question.ID,
			question.Text,
			question.Round,
			question.EventID,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create question")
		}

		questions = append(questions, q)
	}

	return questions, nil
}
