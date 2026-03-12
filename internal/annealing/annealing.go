package annealing

import (
	"fmt"
	"math"
	"math/rand"

	"metopt-lab1/internal/objective"
)

type Config struct {
	MaxIterations      int
	InitialTemperature float64
	MinTemperature     float64
	CoolingRate        float64
	StepSize           float64
	Epsilon            float64
	LowerBounds        []float64
	UpperBounds        []float64
	Seed               int64
}

type Result struct {
	BestPoint        []float64
	BestValue        float64
	Iterations       int
	AcceptedMoves    int
	FinalTemperature float64
}

func DefaultConfig(dimension int) (Config, error) {
	if dimension <= 0 {
		return Config{}, ErrNonPositiveDimension
	}

	lower := make([]float64, dimension)
	upper := make([]float64, dimension)
	for i := 0; i < dimension; i++ {
		lower[i] = -2
		upper[i] = 2
	}

	return Config{
		MaxIterations:      20000,
		InitialTemperature: 10,
		MinTemperature:     1e-6,
		CoolingRate:        0.995,
		StepSize:           0.1,
		Epsilon:            1e-9,
		LowerBounds:        lower,
		UpperBounds:        upper,
		Seed:               1,
	}, nil
}

func Optimize(fn objective.Function, start []float64, cfg Config) (Result, error) {
	if fn == nil {
		return Result{}, ErrNilObjective
	}

	dimension := fn.Dimension()
	if err := validateConfig(cfg, dimension); err != nil {
		return Result{}, err
	}

	if len(start) != dimension {
		return Result{}, ErrStartDimensionMismatch
	}
	if err := validatePointInBounds(start, cfg.LowerBounds, cfg.UpperBounds); err != nil {
		return Result{}, err
	}

	rng := rand.New(rand.NewSource(cfg.Seed))

	currentPoint := copyPoint(start)
	currentValue, err := evaluate(fn, currentPoint, cfg.Epsilon)
	if err != nil {
		return Result{}, err
	}

	bestPoint := copyPoint(currentPoint)
	bestValue := currentValue

	temperature := cfg.InitialTemperature
	acceptedMoves := 0
	iterations := 0

	for i := 0; i < cfg.MaxIterations && temperature > cfg.MinTemperature; i++ {
		candidatePoint := generateNeighbor(currentPoint, cfg.StepSize, cfg.LowerBounds, cfg.UpperBounds, rng)
		candidateValue, err := evaluate(fn, candidatePoint, cfg.Epsilon)
		if err != nil {
			return Result{}, err
		}

		if shouldAccept(candidateValue-currentValue, temperature, rng) {
			currentPoint = candidatePoint
			currentValue = candidateValue
			acceptedMoves++

			if currentValue < bestValue {
				bestValue = currentValue
				bestPoint = copyPoint(currentPoint)
			}
		}

		temperature *= cfg.CoolingRate
		iterations++
	}

	return Result{
		BestPoint:        bestPoint,
		BestValue:        bestValue,
		Iterations:       iterations,
		AcceptedMoves:    acceptedMoves,
		FinalTemperature: temperature,
	}, nil
}

func validateConfig(cfg Config, dimension int) error {
	if cfg.MaxIterations <= 0 {
		return ErrInvalidIterations
	}
	if cfg.InitialTemperature <= 0 {
		return ErrInvalidInitialTemperature
	}
	if cfg.MinTemperature <= 0 {
		return ErrInvalidMinTemperature
	}
	if cfg.MinTemperature >= cfg.InitialTemperature {
		return fmt.Errorf("%w: должна быть меньше начальной температуры", ErrInvalidMinTemperature)
	}
	if cfg.CoolingRate <= 0 || cfg.CoolingRate >= 1 {
		return ErrInvalidCoolingRate
	}
	if cfg.StepSize <= 0 {
		return ErrInvalidStepSize
	}
	if cfg.Epsilon <= 0 {
		return ErrInvalidEpsilon
	}

	if len(cfg.LowerBounds) != dimension || len(cfg.UpperBounds) != dimension {
		return ErrInvalidBounds
	}
	for i := 0; i < dimension; i++ {
		if cfg.LowerBounds[i] >= cfg.UpperBounds[i] {
			return ErrInvalidBounds
		}
	}

	return nil
}

func validatePointInBounds(point, lowerBounds, upperBounds []float64) error {
	for i := 0; i < len(point); i++ {
		if point[i] < lowerBounds[i] || point[i] > upperBounds[i] {
			return ErrStartOutOfBounds
		}
	}
	return nil
}

func evaluate(fn objective.Function, point []float64, eps float64) (float64, error) {
	value, err := fn.Evaluate(point)
	if err != nil {
		return 0, err
	}
	return value.Approx(eps)
}

func generateNeighbor(point []float64, stepSize float64, lowerBounds []float64, upperBounds []float64, rng *rand.Rand) []float64 {
	neighbor := make([]float64, len(point))
	for i := 0; i < len(point); i++ {
		shift := (rng.Float64()*2 - 1) * stepSize
		neighbor[i] = clamp(point[i]+shift, lowerBounds[i], upperBounds[i])
	}
	return neighbor
}

func shouldAccept(delta, temperature float64, rng *rand.Rand) bool {
	if delta <= 0 {
		return true
	}
	probability := math.Exp(-delta / temperature)
	return rng.Float64() < probability
}

func copyPoint(point []float64) []float64 {
	cp := make([]float64, len(point))
	copy(cp, point)
	return cp
}

func clamp(value, minValue, maxValue float64) float64 {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}
