package constructive

import (
	"math"
	"testing"
)

func TestNewConstantApprox(t *testing.T) {
	r := NewConstant(3.14)

	got, err := r.Approx(1e-9)
	if err != nil {
		t.Fatalf("Approx() returned error: %v", err)
	}

	want := 3.14
	if got != want {
		t.Fatalf("Approx() = %v, want %v", got, want)
	}
}

func TestNewConstantApproxInvalidEpsilon(t *testing.T) {
	r := NewConstant(10)

	_, err := r.Approx(0)
	if err == nil {
		t.Fatal("Approx() error = nil, want non-nil")
	}

	if err != ErrNonPositiveEpsilon {
		t.Fatalf("Approx() error = %v, want %v", err, ErrNonPositiveEpsilon)
	}
}

func TestNewFromFunc(t *testing.T) {
	r := NewFromFunc(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrNonPositiveEpsilon
		}

		// Приближаем 1/3 с шагом eps.
		value := 1.0 / 3.0
		approx := math.Round(value/eps) * eps
		return approx, nil
	})

	got1, err := r.Approx(0.1)
	if err != nil {
		t.Fatalf("Approx(0.1) returned error: %v", err)
	}

	got2, err := r.Approx(0.01)
	if err != nil {
		t.Fatalf("Approx(0.01) returned error: %v", err)
	}

	got3, err := r.Approx(0.001)
	if err != nil {
		t.Fatalf("Approx(0.001) returned error: %v", err)
	}

	if math.Abs(got1-0.3) > 1e-12 {
		t.Fatalf("Approx(0.1) = %v, want approximately %v", got1, 0.3)
	}

	if math.Abs(got2-0.33) > 1e-12 {
		t.Fatalf("Approx(0.01) = %v, want approximately %v", got2, 0.33)
	}

	if math.Abs(got3-0.333) > 1e-12 {
		t.Fatalf("Approx(0.001) = %v, want approximately %v", got3, 0.333)
	}
}

func TestNewFromFuncInvalidEpsilon(t *testing.T) {
	r := NewFromFunc(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrNonPositiveEpsilon
		}
		return math.Pi, nil
	})

	_, err := r.Approx(-1)
	if err == nil {
		t.Fatal("Approx() error = nil, want non-nil")
	}

	if err != ErrNonPositiveEpsilon {
		t.Fatalf("Approx() error = %v, want %v", err, ErrNonPositiveEpsilon)
	}
}

func TestNewFromFuncImprovesApproximationWithSmallerEpsilon(t *testing.T) {
	r := NewFromFunc(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrNonPositiveEpsilon
		}

		value := math.Sqrt(2)
		approx := math.Round(value/eps) * eps
		return approx, nil
	})

	coarse, err := r.Approx(0.1)
	if err != nil {
		t.Fatalf("Approx(0.1) returned error: %v", err)
	}

	fine, err := r.Approx(0.001)
	if err != nil {
		t.Fatalf("Approx(0.001) returned error: %v", err)
	}

	target := math.Sqrt(2)

	coarseError := math.Abs(coarse - target)
	fineError := math.Abs(fine - target)

	if fineError > coarseError {
		t.Fatalf("finer approximation is worse: coarse error = %v, fine error = %v", coarseError, fineError)
	}
}
