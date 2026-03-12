package main

import (
	"errors"
	"math"
	"testing"

	"metopt-lab1/internal/annealing"
	"metopt-lab1/internal/objective"
)

func TestEvaluate(t *testing.T) {
	fn := objective.NewRosenbrock2D()

	got, err := evaluate(fn, []float64{1, 1}, 1e-9)
	if err != nil {
		t.Fatalf("evaluate() returned error: %v", err)
	}

	if math.Abs(got) > 1e-12 {
		t.Fatalf("evaluate() = %v, want 0", got)
	}
}

func TestEvaluateDimensionMismatch(t *testing.T) {
	fn := objective.NewRosenbrock2D()

	_, err := evaluate(fn, []float64{1}, 1e-9)
	if !errors.Is(err, objective.ErrPointDimensionMismatch) {
		t.Fatalf("evaluate() error = %v, want %v", err, objective.ErrPointDimensionMismatch)
	}
}

func TestDistanceToOnes(t *testing.T) {
	got := distanceToOnes([]float64{1, 1})
	if got != 0 {
		t.Fatalf("distanceToOnes([1,1]) = %v, want 0", got)
	}
}

func TestDistanceToOnesNonZero(t *testing.T) {
	got := distanceToOnes([]float64{0, -1})
	want := math.Sqrt(5)

	if math.Abs(got-want) > 1e-12 {
		t.Fatalf("distanceToOnes([0,-1]) = %v, want %v", got, want)
	}
}

func TestRunExperiment(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	cfg, err := annealing.DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}
	cfg.MaxIterations = 1000
	cfg.CoolingRate = 0.99
	cfg.Seed = 1

	err = runExperiment("test", fn, []float64{-1.2, 1.0}, cfg)
	if err != nil {
		t.Fatalf("runExperiment() returned error: %v", err)
	}
}

func TestRunExperimentInvalidStart(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	cfg, err := annealing.DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	err = runExperiment("test", fn, []float64{1}, cfg)
	if !errors.Is(err, objective.ErrPointDimensionMismatch) {
		t.Fatalf("runExperiment() error = %v, want %v", err, objective.ErrPointDimensionMismatch)
	}
}

func TestRunRosenbrock2D(t *testing.T) {
	if err := runRosenbrock2D(); err != nil {
		t.Fatalf("runRosenbrock2D() returned error: %v", err)
	}
}

func TestRunRosenbrockND(t *testing.T) {
	if err := runRosenbrockND(3); err != nil {
		t.Fatalf("runRosenbrockND(3) returned error: %v", err)
	}
}

func TestRunRosenbrockNDInvalidDimension(t *testing.T) {
	err := runRosenbrockND(1)
	if !errors.Is(err, objective.ErrInvalidDimension) {
		t.Fatalf("runRosenbrockND(1) error = %v, want %v", err, objective.ErrInvalidDimension)
	}
}
