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

// Find implements the mygame.QuestionRepository repository.
func (r *questionRepository) Find(ctx context.Context, questionID mygame.QuestionID) (mygame.Question, error) {
	questionSQLStruct := sqlbuilder.NewStruct(new(sqlQuestion))

	sb := questionSQLStruct.SelectFrom(sqlQuestionTable)
	sb.Where(sb.E("id", questionID.String()))

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var question sqlQuestion
	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(questionSQLStruct.Addr(&question)...); err != nil {
		if err == sql.ErrNoRows {
			return mygame.Question{}, errors.WrapNotFound(err, "question with %s id not found", questionID.String())
		}

		return mygame.Question{}, errors.Wrap(err, "failed to find question")
	}

	q, err := mygame.NewQuestion(
		question.ID,
		question.Text,
	)
	if err != nil {
		return mygame.Question{}, err
	}

	return q, nil
}
