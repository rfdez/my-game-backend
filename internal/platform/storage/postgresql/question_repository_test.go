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

func Test_QuestionRepository_Find_RepositoryError(t *testing.T) {
	questionIDVO, err := mygame.NewQuestionID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		`SELECT questions.id, questions.text FROM questions WHERE id = $1`).
		WithArgs(questionIDVO.String()).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewQuestionRepository(db, 1*time.Millisecond)

	_, err = repo.Find(context.Background(), questionIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_QuestionRepository_Find_Succeed(t *testing.T) {
	questionIDVO, err := mygame.NewQuestionID("1b334c46-9792-4603-b16e-b3e05146b778")
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "text"}).
		AddRow("1b334c46-9792-4603-b16e-b3e05146b778", "Question 1")

	sqlMock.ExpectQuery(
		`SELECT questions.id, questions.text FROM questions WHERE id = $1`).
		WithArgs(questionIDVO.String()).
		WillReturnRows(rows)

	repo := postgresql.NewQuestionRepository(db, 1*time.Millisecond)

	question, err := repo.Find(context.Background(), questionIDVO)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, "1b334c46-9792-4603-b16e-b3e05146b778", question.ID().String())
}
