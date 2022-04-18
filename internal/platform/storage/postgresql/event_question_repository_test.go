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

func Test_EventQuestionRepository_SearchByEventID_RepositoryError(t *testing.T) {
	eventIDVO, err := mygame.NewEventID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		`SELECT event_questions.event_id, event_questions.question_id, event_questions.round FROM event_questions WHERE event_id = $1`).
		WithArgs(eventIDVO.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewEventQuestionRepository(db, 1*time.Millisecond)

	_, err = repo.SearchByEventID(context.Background(), eventIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_EventQuestionRepository_SearchByEventID_Succeed(t *testing.T) {
	eventIDVO, err := mygame.NewEventID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"event_id", "question_id", "round"}).
		AddRow("1b334c46-9792-4603-b16e-b3e05146b778", "62bac31e-d6ba-4143-a456-637a65ca62e9", 1).
		AddRow("1b334c46-9792-4603-b16e-b3e05146b778", "e55917ea-2cc7-44cc-bc10-4086e5cb9fea", 1)

	sqlMock.ExpectQuery(
		`SELECT event_questions.event_id, event_questions.question_id, event_questions.round FROM event_questions WHERE event_id = $1`).
		WithArgs(eventIDVO.String()).
		WillReturnRows(rows)

	repo := postgresql.NewEventQuestionRepository(db, 1*time.Millisecond)

	events, err := repo.SearchByEventID(context.Background(), eventIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))
}
