package db

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestAddToHistoryExecutesInsert(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	mock.ExpectExec("INSERT INTO arith_history").
		WithArgs(sqlmock.AnyArg(), int32(7), "addition").
		WillReturnResult(sqlmock.NewResult(1, 1))

	adapter := Adapter{db: sqlDB}

	err = adapter.AddToHistory(7, "addition")
	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAddToHistoryPropagatesExecError(t *testing.T) {
	sqlDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer sqlDB.Close()

	execErr := errors.New("insert failed")
	mock.ExpectExec("INSERT INTO arith_history").
		WithArgs(sqlmock.AnyArg(), int32(3), "division").
		WillReturnError(execErr)

	adapter := Adapter{db: sqlDB}

	err = adapter.AddToHistory(3, "division")
	require.Error(t, err)
	require.True(t, errors.Is(err, execErr))
	require.NoError(t, mock.ExpectationsWereMet())
}

