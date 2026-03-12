package annealing

import (
	"errors"
	"math/rand"
	"testing"

	"metopt-lab1/internal/constructive"
	"metopt-lab1/internal/objective"
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

func TestDefaultConfigValid(t *testing.T) {
	cfg, err := DefaultConfig(3)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	if len(cfg.LowerBounds) != 3 || len(cfg.UpperBounds) != 3 {
		t.Fatalf("bounds dimensions = (%d, %d), want (3, 3)", len(cfg.LowerBounds), len(cfg.UpperBounds))
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

func TestOptimizeNilObjective(t *testing.T) {
	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	_, err = Optimize(nil, []float64{0, 0}, cfg)
	if !errors.Is(err, ErrNilObjective) {
		t.Fatalf("Optimize() error = %v, want %v", err, ErrNilObjective)
	}
}

func TestOptimizeStartDimensionMismatch(t *testing.T) {
	fn := sphereObjective{dimension: 2}
	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	_, err = Optimize(fn, []float64{0}, cfg)
	if !errors.Is(err, ErrStartDimensionMismatch) {
		t.Fatalf("Optimize() error = %v, want %v", err, ErrStartDimensionMismatch)
	}
}

func TestOptimizeStartOutOfBounds(t *testing.T) {
	fn := sphereObjective{dimension: 2}
	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	_, err = Optimize(fn, []float64{10, 0}, cfg)
	if !errors.Is(err, ErrStartOutOfBounds) {
		t.Fatalf("Optimize() error = %v, want %v", err, ErrStartOutOfBounds)
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

func TestValidateConfigValid(t *testing.T) {
	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}

	if err := validateConfig(cfg, 2); err != nil {
		t.Fatalf("validateConfig() returned error: %v", err)
	}
}

func TestValidateConfigInvalidBounds(t *testing.T) {
	cfg, err := DefaultConfig(2)
	if err != nil {
		t.Fatalf("DefaultConfig() returned error: %v", err)
	}
	cfg.UpperBounds = []float64{2}

	err = validateConfig(cfg, 2)
	if !errors.Is(err, ErrInvalidBounds) {
		t.Fatalf("validateConfig() error = %v, want %v", err, ErrInvalidBounds)
	}
}

func TestValidatePointInBounds(t *testing.T) {
	err := validatePointInBounds([]float64{0, 0}, []float64{-1, -1}, []float64{1, 1})
	if err != nil {
		t.Fatalf("validatePointInBounds() returned error: %v", err)
	}
}

func TestValidatePointInBoundsOutOfBounds(t *testing.T) {
	err := validatePointInBounds([]float64{2, 0}, []float64{-1, -1}, []float64{1, 1})
	if !errors.Is(err, ErrStartOutOfBounds) {
		t.Fatalf("validatePointInBounds() error = %v, want %v", err, ErrStartOutOfBounds)
	}
}

func TestEvaluate(t *testing.T) {
	fn := sphereObjective{dimension: 2}

	got, err := evaluate(fn, []float64{3, 4}, 1e-9)
	if err != nil {
		t.Fatalf("evaluate() returned error: %v", err)
	}

	if got != 25 {
		t.Fatalf("evaluate() = %v, want 25", got)
	}
}

func TestEvaluateInvalidPoint(t *testing.T) {
	fn := sphereObjective{dimension: 2}

	_, err := evaluate(fn, []float64{3}, 1e-9)
	if !errors.Is(err, objective.ErrPointDimensionMismatch) {
		t.Fatalf("evaluate() error = %v, want %v", err, objective.ErrPointDimensionMismatch)
	}
}

func TestGenerateNeighborRespectsBounds(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	point := []float64{0.5, -0.5}

	neighbor := generateNeighbor(point, 1.0, []float64{0, -1}, []float64{1, 0}, rng)

	if neighbor[0] < 0 || neighbor[0] > 1 || neighbor[1] < -1 || neighbor[1] > 0 {
		t.Fatalf("neighbor %v is out of bounds", neighbor)
	}
}

func TestGenerateNeighborWithZeroStepSizeKeepsPoint(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	point := []float64{0.5, -0.5}

	neighbor := generateNeighbor(point, 0, []float64{-1, -1}, []float64{1, 1}, rng)

	if neighbor[0] != point[0] || neighbor[1] != point[1] {
		t.Fatalf("neighbor = %v, want %v", neighbor, point)
	}
}

func TestShouldAcceptForBetterPoint(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	if !shouldAccept(-1, 1, rng) {
		t.Fatal("shouldAccept() = false, want true for improving move")
	}
}

func TestShouldAcceptRejectsHugeWorseMove(t *testing.T) {
	rng := rand.New(rand.NewSource(1))

	if shouldAccept(1e9, 1, rng) {
		t.Fatal("shouldAccept() = true, want false for near-zero acceptance probability")
	}
}

func TestCopyPointCreatesIndependentCopy(t *testing.T) {
	point := []float64{1, 2}
	cp := copyPoint(point)
	point[0] = 99

	if cp[0] != 1 {
		t.Fatalf("copyPoint() shares storage with source, copy = %v", cp)
	}
}

func TestCopyPointPreservesValues(t *testing.T) {
	point := []float64{3, 4}
	cp := copyPoint(point)

	if cp[0] != point[0] || cp[1] != point[1] {
		t.Fatalf("copyPoint() = %v, want %v", cp, point)
	}
}

func TestClampInsideRange(t *testing.T) {
	got := clamp(0.5, 0, 1)
	if got != 0.5 {
		t.Fatalf("clamp() = %v, want 0.5", got)
	}
}

func TestClampOutsideRange(t *testing.T) {
	got := clamp(2, 0, 1)
	if got != 1 {
		t.Fatalf("clamp() = %v, want 1", got)
	}
}
