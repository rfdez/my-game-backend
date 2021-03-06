package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
)

type answerRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewAnswerRepository initializes a PostgreSQL implementation of mygame.AnswerRepository.
func NewAnswerRepository(db *sql.DB, dbTimeout time.Duration) *answerRepository {
	return &answerRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// FindByEventIDAndQuestionID implements the mygame.AnswerRepository repository.
func (r *answerRepository) FindByEventIDAndQuestionID(ctx context.Context, eventID mygame.EventID, questionID mygame.QuestionID) (mygame.Answer, error) {
	answerSQLStruct := sqlbuilder.NewStruct(new(sqlAnswer))

	sb := answerSQLStruct.SelectFrom(sqlAnswerTable)
	sb.Where(sb.E("event_id", eventID.String()), sb.E("question_id", questionID.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var answer sqlAnswer
	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(answerSQLStruct.Addr(&answer)...); err != nil {
		if err == sql.ErrNoRows {
			return mygame.Answer{}, errors.NewNotFound("answer with %s event id and %s question id not found", eventID.String(), questionID.String())
		}

		return mygame.Answer{}, errors.Wrap(err, "failed to search answer")
	}

	ans, err := mygame.NewAnswer(
		answer.EventID,
		answer.QuestionID,
		answer.Text,
	)
	if err != nil {
		return mygame.Answer{}, err
	}

	return ans, nil
}
