package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := New()

	answer, err := arith.Addition(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(2))
}

func TestSubtraction(t *testing.T) {
	arith := New()

	answer, err := arith.Subtraction(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(0))
}

func TestMultiplication(t *testing.T) {
	arith := New()

	answer, err := arith.Multiplication(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestDivision(t *testing.T) {
	arith := New()

	answer, err := arith.Division(1, 1)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(1))
}

func TestAdditionWithNegativeNumbers(t *testing.T) {
	arith := New()

	answer, err := arith.Addition(-4, -6)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(-10))
}

func TestMultiplicationWithNegativeNumbers(t *testing.T) {
	arith := New()

	answer, err := arith.Multiplication(-3, 5)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(-15))
}

func TestDivisionByNegativeNumber(t *testing.T) {
	arith := New()

	answer, err := arith.Division(9, -3)
	if err != nil {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, answer, int32(-3))
}

func TestDivisionByZeroPanics(t *testing.T) {
	arith := New()

	require.Panics(t, func() {
		_, _ = arith.Division(1, 0)
	})
}
