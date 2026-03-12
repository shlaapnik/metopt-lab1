package objective

import (
	"errors"
	"math"
	"testing"
)

func TestRosenbrock2DMinimum(t *testing.T) {
	fn := NewRosenbrock2D()

	result, err := fn.Evaluate([]float64{1, 1})
	if err != nil {
		t.Fatalf("Evaluate() returned error: %v", err)
	}

	got, err := result.Approx(1e-9)
	if err != nil {
		t.Fatalf("Approx() returned error: %v", err)
	}

	if math.Abs(got) > 1e-12 {
		t.Fatalf("Evaluate([1,1]) = %v, want 0", got)
	}
}

func TestRosenbrock2DKnownPoint(t *testing.T) {
	fn := NewRosenbrock2D()

	result, err := fn.Evaluate([]float64{0, 0})
	if err != nil {
		t.Fatalf("Evaluate() returned error: %v", err)
	}

	got, err := result.Approx(1e-9)
	if err != nil {
		t.Fatalf("Approx() returned error: %v", err)
	}

	want := 1.0
	if math.Abs(got-want) > 1e-12 {
		t.Fatalf("Evaluate([0,0]) = %v, want %v", got, want)
	}
}

func TestRosenbrockNDimMinimum(t *testing.T) {
	fn, err := NewRosenbrock(5)
	if err != nil {
		t.Fatalf("NewRosenbrock() returned error: %v", err)
	}

	result, err := fn.Evaluate([]float64{1, 1, 1, 1, 1})
	if err != nil {
		t.Fatalf("Evaluate() returned error: %v", err)
	}

	got, err := result.Approx(1e-9)
	if err != nil {
		t.Fatalf("Approx() returned error: %v", err)
	}

	if math.Abs(got) > 1e-12 {
		t.Fatalf("Evaluate([1,...,1]) = %v, want 0", got)
	}
}

func TestRosenbrockInvalidDimension(t *testing.T) {
	_, err := NewRosenbrock(1)
	if !errors.Is(err, ErrInvalidDimension) {
		t.Fatalf("NewRosenbrock(1) error = %v, want %v", err, ErrInvalidDimension)
	}
}

func TestRosenbrockPointDimensionMismatch(t *testing.T) {
	fn := NewRosenbrock2D()

	_, err := fn.Evaluate([]float64{1, 2, 3})
	if !errors.Is(err, ErrPointDimensionMismatch) {
		t.Fatalf("Evaluate() error = %v, want %v", err, ErrPointDimensionMismatch)
	}
}
