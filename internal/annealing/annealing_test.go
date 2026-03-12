package annealing

import (
	"errors"
	"testing"

	"github.com/shlaapnik/metopt-lab1/internal/constructive"
	"github.com/shlaapnik/metopt-lab1/internal/objective"
)

type sphereObjective struct {
	dimension int
}

func (s sphereObjective) Dimension() int {
	return s.dimension
}

func (s sphereObjective) Evaluate(point []float64) (constructive.Real, error) {
	if len(point) != s.dimension {
		return nil, objective.ErrPointDimensionMismatch
	}

	value := 0.0
	for _, coordinate := range point {
		value += coordinate * coordinate
	}
	return constructive.NewConstant(value), nil
}

func TestDefaultConfigInvalidDimension(t *testing.T) {
	_, err := DefaultConfig(0)
	if !errors.Is(err, ErrNonPositiveDimension) {
		t.Fatalf("DefaultConfig(0) error = %v, want %v", err, ErrNonPositiveDimension)
	}
}

func TestOptimizeInvalidConfig(t *testing.T) {
	fn := sphereObjective{dimension: 2}
	start := []float64{1, 1}

	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}
	cfg.MaxIterations = 0

	_, err = Optimize(fn, start, cfg)
	if !errors.Is(err, ErrInvalidIterations) {
		t.Fatalf("Optimize() error = %v, want %v", err, ErrInvalidIterations)
	}
}

func TestOptimizeImprovesSphereObjective(t *testing.T) {
	fn := sphereObjective{dimension: 2}
	start := []float64{3, -4}

	cfg := Config{
		MaxIterations:      15000,
		InitialTemperature: 10,
		MinTemperature:     1e-6,
		CoolingRate:        0.995,
		StepSize:           0.3,
		Epsilon:            1e-9,
		LowerBounds:        []float64{-5, -5},
		UpperBounds:        []float64{5, 5},
		Seed:               42,
	}

	result, err := Optimize(fn, start, cfg)
	if err != nil {
		t.Fatalf("Optimize() returned error: %v", err)
	}

	startValue := 25.0
	if result.BestValue >= startValue {
		t.Fatalf("BestValue = %v, want less than start value %v", result.BestValue, startValue)
	}

	if result.BestValue > 0.5 {
		t.Fatalf("BestValue = %v, want <= 0.5", result.BestValue)
	}

	if len(result.BestPoint) != 2 {
		t.Fatalf("BestPoint dimension = %d, want 2", len(result.BestPoint))
	}
}

func TestOptimizeRosenbrock2D(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	start := []float64{-1.2, 1.0}

	cfg := Config{
		MaxIterations:      30000,
		InitialTemperature: 8,
		MinTemperature:     1e-6,
		CoolingRate:        0.999,
		StepSize:           0.2,
		Epsilon:            1e-9,
		LowerBounds:        []float64{-2, -2},
		UpperBounds:        []float64{2, 2},
		Seed:               7,
	}

	result, err := Optimize(fn, start, cfg)
	if err != nil {
		t.Fatalf("Optimize() returned error: %v", err)
	}

	startEval, err := fn.Evaluate(start)
	if err != nil {
		t.Fatalf("Evaluate(start) returned error: %v", err)
	}
	startValue, err := startEval.Approx(1e-9)
	if err != nil {
		t.Fatalf("Approx(start) returned error: %v", err)
	}

	if result.BestValue >= startValue {
		t.Fatalf("BestValue = %v, want less than start value %v", result.BestValue, startValue)
	}

	if result.BestValue >= 5 {
		t.Fatalf("BestValue = %v, want < 5", result.BestValue)
	}
}
