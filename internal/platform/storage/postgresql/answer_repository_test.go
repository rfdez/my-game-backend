package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/platform/storage/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_AnswerRepository_FindByEventIDAndQuestionID_RepositoryError(t *testing.T) {
	eventIDVO, err := mygame.NewEventID("2dfb644e-1dbd-4551-ab2c-5d8a2555bcf8")
	require.NoError(t, err)

	questionIDVO, err := mygame.NewQuestionID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		`SELECT answers.event_id, answers.question_id, answers.text FROM answers WHERE event_id = $1 AND question_id = $2`).
		WithArgs(eventIDVO.String(), questionIDVO.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewAnswerRepository(db, 1*time.Millisecond)

	_, err = repo.FindByEventIDAndQuestionID(context.Background(), eventIDVO, questionIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_AnswerRepository_FindByEventIDAndQuestionID_Succeed(t *testing.T) {
	eventIDVO, err := mygame.NewEventID("2dfb644e-1dbd-4551-ab2c-5d8a2555bcf8")
	require.NoError(t, err)

	questionIDVO, err := mygame.NewQuestionID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"event_id", "question_id", "text"}).
		AddRow("2dfb644e-1dbd-4551-ab2c-5d8a2555bcf8", "1b334c46-9792-4603-b16e-b3e05146b778", "Answer 1")

	sqlMock.ExpectQuery(
		`SELECT answers.event_id, answers.question_id, answers.text FROM answers WHERE event_id = $1 AND question_id = $2`).
		WithArgs(eventIDVO.String(), questionIDVO.String()).
		WillReturnRows(rows)

	repo := postgresql.NewAnswerRepository(db, 1*time.Millisecond)

	answer, err := repo.FindByEventIDAndQuestionID(context.Background(), eventIDVO, questionIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, "2dfb644e-1dbd-4551-ab2c-5d8a2555bcf8", answer.EventID().String())
	assert.Equal(t, "1b334c46-9792-4603-b16e-b3e05146b778", answer.QuestionID().String())
}
