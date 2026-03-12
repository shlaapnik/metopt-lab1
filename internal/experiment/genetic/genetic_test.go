package experiment

import (
	"math"
	"math/rand"
	"testing"

	"metopt-lab1/internal/genetic"
	"metopt-lab1/internal/genetic/crossover"
	"metopt-lab1/internal/genetic/mutation"
	"metopt-lab1/internal/genetic/replacement"
	"metopt-lab1/internal/genetic/selection"
	"metopt-lab1/internal/objective"
)

func TestRunGeneticPerformance(t *testing.T) {
	fn := objective.NewRosenbrock2D()
	bounds := [][]float64{{-30.0, 30.0}, {-30.0, 30.0}}

	fitness := func(genes []float64) float64 {
		realVal, err := fn.Evaluate(genes)
		if err != nil {
			return math.MaxFloat64
		}

		val, _ := realVal.Approx(1e-9)
		return val
	}

	sel := selection.NewTournament[float64](3)
	cross := crossover.NewSBXCrossover[float64](0.9, 20.0, bounds)
	mut := mutation.NewGaussianMutator[float64](0.1, bounds, 0.01)
	repl := replacement.NewElitismReplacement[float64](8)

	genFunc := func(index int, b []float64) float64 {
		return b[0] + rand.Float64()*(b[1]-b[0])
	}

	engine := genetic.NewEngine[float64](
		bounds,
		fitness,
		genFunc,
		sel,
		cross,
		mut,
		repl,
		350,
		100,
		-1.0,
	)

	run := RunGenetic(engine)

	if run.Duration <= 0 {
		t.Errorf("expected positive duration, got %v", run.Duration)
	}

	t.Logf("GA Performance Test: Rosenbrock 2D")
	t.Logf("GA Rosenbrock finished in %v", run.Duration)
	t.Logf("Best Fitness: %f", run.Best.Fitness)
	t.Logf("Best Genes: %v", run.Best.Genes)

	if run.Best.Fitness > 1.0 {
		t.Log("Warning: GA did not converge well, consider tuning parameters")
	}
}

func TestRunGeneticPerformance5D(t *testing.T) {
	dim := 5
	fn, err := objective.NewRosenbrock(dim)
	if err != nil {
		t.Fatalf("failed to create Rosenbrock %dD: %v", dim, err)
	}

	bounds := make([][]float64, dim)
	for i := 0; i < dim; i++ {
		bounds[i] = []float64{-10.0, 10.0}
	}

	fitness := func(genes []float64) float64 {
		realVal, err := fn.Evaluate(genes)
		if err != nil {
			return 1e18
		}
		val, _ := realVal.Approx(1e-9)
		return val
	}

	sel := selection.NewTournament[float64](5)
	cross := crossover.NewSBXCrossover[float64](0.9, 20.0, bounds)
	mut := mutation.NewGaussianMutator[float64](0.2, bounds, 0.02)
	repl := replacement.NewElitismReplacement[float64](10)

	genFunc := func(index int, b []float64) float64 {
		return b[0] + rand.Float64()*(b[1]-b[0])
	}

	engine := genetic.NewEngine[float64](
		bounds,
		fitness,
		genFunc,
		sel,
		cross,
		mut,
		repl,
		400,
		200,
		-1.0,
	)

	run := RunGenetic(engine)

	t.Logf("GA Performance Test: Rosenbrock %dD", dim)
	t.Logf("Duration: %v", run.Duration)
	t.Logf("Best Fitness found: %f", run.Best.Fitness)
	t.Logf("Best Point: %v", run.Best.Genes)

	if run.Duration <= 0 {
		t.Errorf("expected positive duration, got %v", run.Duration)
	}

	if len(run.Best.Genes) != dim {
		t.Errorf("expected %d genes, got %d", dim, len(run.Best.Genes))
	}

	target := 0.0
	diff := math.Abs(run.Best.Fitness - target)
	t.Logf("Difference from global minimum (0.0): %f", diff)
}
