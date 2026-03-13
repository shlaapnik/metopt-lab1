package experiment

import (
	"fmt"
	"math"
	"math/rand"
	"metopt-lab1/internal/annealing"
	expAnn "metopt-lab1/internal/experiment/annealing"
	expGen "metopt-lab1/internal/experiment/genetic"
	"metopt-lab1/internal/genetic"
	"metopt-lab1/internal/genetic/crossover"
	"metopt-lab1/internal/genetic/mutation"
	"metopt-lab1/internal/genetic/replacement"
	"metopt-lab1/internal/genetic/selection"
	"metopt-lab1/internal/objective"
)

func RunAllExperiments() error {
	fmt.Println("=== Сравнение алгоритмов оптимизации на функции Розенброка ===")

	fmt.Println("\n=== Simulated Annealing ===")
	if err := runAnnealing2D(); err != nil {
		return fmt.Errorf("ошибка SA 2D: %w", err)
	}
	if err := runAnnealingND(5); err != nil {
		return fmt.Errorf("ошибка SA 5D: %w", err)
	}

	fmt.Println("\n=== Genetic Algorithm ===")
	if err := runGenetic2D(); err != nil {
		return fmt.Errorf("ошибка GA 2D: %w", err)
	}
	if err := runGeneticND(5); err != nil {
		return fmt.Errorf("ошибка GA 5D: %w", err)
	}

	return nil
}

func runAnnealing2D() error {
	fn := objective.NewRosenbrock2D()
	start := []float64{-1.2, 1.0}

	cfg, err := annealing.DefaultConfig(fn.Dimension())
	if err != nil {
		return err
	}

	cfg.MaxIterations = 30000
	cfg.InitialTemperature = 8
	cfg.CoolingRate = 0.999
	cfg.StepSize = 0.2
	cfg.LowerBounds = []float64{-2, -2}
	cfg.UpperBounds = []float64{2, 2}
	cfg.Seed = 7

	return printAnnealingResult("SA 2D", fn, start, cfg)
}

func runAnnealingND(dimension int) error {
	fn, err := objective.NewRosenbrock(dimension)
	if err != nil {
		return err
	}

	start := make([]float64, dimension)
	start[0] = -1.2
	for i := 1; i < dimension; i++ {
		start[i] = 1.0
	}

	cfg, err := annealing.DefaultConfig(fn.Dimension())
	if err != nil {
		return err
	}

	cfg.MaxIterations = 60000
	cfg.InitialTemperature = 10
	cfg.CoolingRate = 0.9995
	cfg.StepSize = 0.25
	cfg.Seed = 11

	return printAnnealingResult(fmt.Sprintf("SA %dD", dimension), fn, start, cfg)
}

func printAnnealingResult(name string, fn objective.Function, start []float64, cfg annealing.Config) error {
	startValue, err := evaluate(fn, start, cfg.Epsilon)
	if err != nil {
		return err
	}

	run, err := expAnn.RunAnnealing(fn, start, cfg)
	if err != nil {
		return err
	}

	fmt.Printf("\n--- %s ---\n", name)
	fmt.Printf("Стартовая точка:       %v\n", start)
	fmt.Printf("Стартовое значение:    %.8f\n", startValue)
	fmt.Printf("Лучшая точка:          %v\n", run.Result.BestPoint)
	fmt.Printf("Лучшее значение:       %.8f\n", run.Result.BestValue)
	fmt.Printf("Итераций:              %d\n", run.Result.Iterations)
	fmt.Printf("Время выполнения:      %s\n", run.Duration)

	return nil
}

func runGenetic2D() error {
	fn := objective.NewRosenbrock2D()
	bounds := [][]float64{{-2.0, 2.0}, {-2.0, 2.0}}

	engine := buildGeneticEngine(fn, 2, bounds, 200, 150, 42)
	return printGeneticResult("GA 2D", engine)
}

func runGeneticND(dimension int) error {
	fn, err := objective.NewRosenbrock(dimension)
	if err != nil {
		return err
	}

	bounds := make([][]float64, dimension)
	for i := 0; i < dimension; i++ {
		bounds[i] = []float64{-2.0, 2.0}
	}

	engine := buildGeneticEngine(fn, dimension, bounds, 500, 400, 42)
	return printGeneticResult(fmt.Sprintf("GA %dD", dimension), engine)
}

func buildGeneticEngine(fn objective.Function, dim int, bounds [][]float64, popSize, gen, seed int) *genetic.Engine[float64] {
	fitness := func(genes []float64) float64 {
		realVal, err := fn.Evaluate(genes)
		if err != nil {
			return 1e18
		}
		val, _ := realVal.Approx(1e-9)
		return val
	}
	genFunc := func(index int, b []float64) float64 { return b[0] + rand.Float64()*(b[1]-b[0]) }

	sel := selection.NewTournament[float64](3, seed)
	cross := crossover.NewSBXCrossover[float64](0.9, 20.0, bounds, seed)
	mut := mutation.NewGaussianMutator[float64](0.1, bounds, 0.05, seed)
	repl := replacement.NewElitismReplacement[float64](10)

	return genetic.NewEngine[float64](bounds, fitness, genFunc, sel, cross, mut, repl, popSize, gen, -1.0, seed)
}

func printGeneticResult(name string, e *genetic.Engine[float64]) error {
	run := expGen.RunGenetic(e)

	fmt.Printf("\n--- %s ---\n", name)
	fmt.Printf("Размер популяции:      %d\n", e.PopSize)
	fmt.Printf("Поколений:             %d\n", e.Generations)
	fmt.Printf("Лучшая точка:          %v\n", run.Best.Genes)
	fmt.Printf("Лучшее значение:       %.8f\n", run.Best.Fitness)
	fmt.Printf("Расстояние до [1..1]:  %.8f\n", distanceToOnes(run.Best.Genes))
	fmt.Printf("Время выполнения:      %s\n", run.Duration)

	return nil
}

func evaluate(fn objective.Function, point []float64, eps float64) (float64, error) {
	value, err := fn.Evaluate(point)
	if err != nil {
		return 0, err
	}
	return value.Approx(eps)
}

func distanceToOnes(point []float64) float64 {
	sum := 0.0
	for _, x := range point {
		diff := x - 1.0
		sum += diff * diff
	}
	return math.Sqrt(sum)
}
