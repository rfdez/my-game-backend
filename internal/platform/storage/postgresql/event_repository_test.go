package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/platform/storage/postgresql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_EventRepository_SearchAll_RepositoryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectQuery(
		`SELECT events.id, events.name, events.date, events.keywords FROM events`).
		WillReturnError(errors.New("error"))

	repo := postgresql.NewEventRepository(db, 1*time.Millisecond)

	_, err = repo.SearchAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_EventRepository_SearchAll_Succeed(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	rows := sqlmock.NewRows([]string{"id", "name", "date", "keywords"}).
		AddRow("1b334c46-9792-4603-b16e-b3e05146b778", "event1", time.Now(), "{keyword1,keyword2}").
		AddRow("57450d62-79df-41b6-a4b1-5ee04eed7188", "event2", time.Now(), "{keyword3,keyword4}")

	sqlMock.ExpectQuery(
		`SELECT events.id, events.name, events.date, events.keywords FROM events`).
		WillReturnRows(rows)

	repo := postgresql.NewEventRepository(db, 1*time.Millisecond)

	events, err := repo.SearchAll(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.Equal(t, 2, len(events))
}
