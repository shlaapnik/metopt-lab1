package experiment

import (
	"errors"
	"testing"

	"metopt-lab1/internal/annealing"
	"metopt-lab1/internal/objective"
)

func TestRunAnnealing(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	start := []float64{-1.2, 1.0}

	cfg := annealing.Config{
		MaxIterations:      2000,
		InitialTemperature: 8,
		MinTemperature:     1e-6,
		CoolingRate:        0.995,
		StepSize:           0.2,
		Epsilon:            1e-9,
		LowerBounds:        []float64{-2, -2},
		UpperBounds:        []float64{2, 2},
		Seed:               5,
	}

	run, err := RunAnnealing(fn, start, cfg)
	if err != nil {
		t.Fatalf("RunAnnealing() returned error: %v", err)
	}

	if run.Duration <= 0 {
		t.Fatalf("Duration = %v, want positive", run.Duration)
	}

	if len(run.Result.BestPoint) != 2 {
		t.Fatalf("BestPoint dimension = %d, want 2", len(run.Result.BestPoint))
	}
}

func TestRunAnnealingInvalidConfig(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	start := []float64{-1.2, 1.0}

	cfg, err := annealing.DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}
	cfg.MaxIterations = 0

	_, err = RunAnnealing(fn, start, cfg)
	if !errors.Is(err, annealing.ErrInvalidIterations) {
		t.Fatalf("RunAnnealing() error = %v, want %v", err, annealing.ErrInvalidIterations)
	}
}
