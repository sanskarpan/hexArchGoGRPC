package api

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type spyDbPort struct {
	calls []historyCall
	err   error
}

type historyCall struct {
	answer    int32
	operation string
}

func (s *spyDbPort) CloseDbConnection() {}

func (s *spyDbPort) AddToHistory(answer int32, operation string) error {
	s.calls = append(s.calls, historyCall{answer: answer, operation: operation})
	return s.err
}

type stubArithmetic struct {
	additionResult       int32
	additionErr          error
	subtractionResult    int32
	subtractionErr       error
	multiplicationResult int32
	multiplicationErr    error
	divisionResult       int32
	divisionErr          error
}

func (s stubArithmetic) Addition(a int32, b int32) (int32, error) {
	if s.additionErr != nil {
		return 0, s.additionErr
	}
	return s.additionResult, nil
}

func (s stubArithmetic) Subtraction(a int32, b int32) (int32, error) {
	if s.subtractionErr != nil {
		return 0, s.subtractionErr
	}
	return s.subtractionResult, nil
}

func (s stubArithmetic) Multiplication(a int32, b int32) (int32, error) {
	if s.multiplicationErr != nil {
		return 0, s.multiplicationErr
	}
	return s.multiplicationResult, nil
}

func (s stubArithmetic) Division(a int32, b int32) (int32, error) {
	if s.divisionErr != nil {
		return 0, s.divisionErr
	}
	return s.divisionResult, nil
}

func TestApplicationOperationsPersistHistory(t *testing.T) {
	spyDb := &spyDbPort{}
	app := NewApplication(spyDb, stubArithmetic{
		additionResult:       7,
		subtractionResult:    -3,
		multiplicationResult: 20,
		divisionResult:       2,
	})

	tests := []struct {
		name              string
		invoke            func(*Application) (int32, error)
		expectedResult    int32
		expectedOperation string
	}{
		{
			name: "addition",
			invoke: func(app *Application) (int32, error) {
				return app.GetAddition(3, 4)
			},
			expectedResult:    7,
			expectedOperation: "addition",
		},
		{
			name: "subtraction",
			invoke: func(app *Application) (int32, error) {
				return app.GetSubtraction(7, 10)
			},
			expectedResult:    -3,
			expectedOperation: "subtraction",
		},
		{
			name: "multiplication",
			invoke: func(app *Application) (int32, error) {
				return app.GetMultiplication(4, 5)
			},
			expectedResult:    20,
			expectedOperation: "multiplication",
		},
		{
			name: "division",
			invoke: func(app *Application) (int32, error) {
				return app.GetDivision(8, 4)
			},
			expectedResult:    2,
			expectedOperation: "division",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spyDb.calls = nil
			result, err := tt.invoke(app)
			require.NoError(t, err)
			require.Equal(t, tt.expectedResult, result)
			require.Len(t, spyDb.calls, 1)
			require.Equal(t, historyCall{answer: tt.expectedResult, operation: tt.expectedOperation}, spyDb.calls[0])
		})
	}
}

func TestApplicationReturnsErrorWhenArithmeticFails(t *testing.T) {
	expectedErr := errors.New("arith failure")
	spyDb := &spyDbPort{}
	app := NewApplication(spyDb, stubArithmetic{additionErr: expectedErr})

	result, err := app.GetAddition(1, 2)
	require.Error(t, err)
	require.True(t, errors.Is(err, expectedErr))
	require.Zero(t, result)
	require.Empty(t, spyDb.calls)
}

func TestApplicationReturnsErrorWhenPersistingHistoryFails(t *testing.T) {
	expectedErr := errors.New("history failure")
	spyDb := &spyDbPort{err: expectedErr}
	app := NewApplication(spyDb, stubArithmetic{additionResult: 3})

	result, err := app.GetAddition(1, 2)
	require.Error(t, err)
	require.True(t, errors.Is(err, expectedErr))
	require.Zero(t, result)
	require.Len(t, spyDb.calls, 1)
	require.Equal(t, historyCall{answer: 3, operation: "addition"}, spyDb.calls[0])
}

