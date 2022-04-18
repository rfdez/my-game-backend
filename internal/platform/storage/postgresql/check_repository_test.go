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

func Test_CheckRepository_Status_RepositoryError(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	sqlMock.ExpectPing().WillReturnError(errors.New("error"))

	repo := postgresql.NewCheckRepository(db, 1*time.Millisecond)

	err = repo.Status(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func Test_CheckRepository_Status_Succeed(t *testing.T) {
	db, sqlMock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))
	require.NoError(t, err)

	sqlMock.ExpectPing().WillReturnError(nil)

	repo := postgresql.NewCheckRepository(db, 1*time.Millisecond)

	err = repo.Status(context.Background())

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
