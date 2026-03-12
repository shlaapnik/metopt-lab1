package main

import (
	"fmt"
	"log"
	"math"

	"metopt-lab1/internal/annealing"
	"metopt-lab1/internal/experiment"
	"metopt-lab1/internal/objective"
)

func main() {
	fmt.Println("=== Эксперименты метода имитации отжига ===")

	if err := runRosenbrock2D(); err != nil {
		log.Fatalf("сбой эксперимента для функции Розенброка 2D: %v", err)
	}

	if err := runRosenbrockND(5); err != nil {
		log.Fatalf("сбой эксперимента для функции Розенброка N-D: %v", err)
	}
}

func runRosenbrock2D() error {
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

	return runExperiment("Функция Розенброка 2D", fn, start, cfg)
}

func runRosenbrockND(dimension int) error {
	fn, err := objective.NewRosenbrock(dimension)
	if err != nil {
		return err
	}

	start := make([]float64, dimension)
	start[0] = -1.2
	for i := 1; i < dimension; i++ {
		start[i] = 1
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

	return runExperiment(fmt.Sprintf("Функция Розенброка %dD", dimension), fn, start, cfg)
}

func runExperiment(name string, fn objective.Function, start []float64, cfg annealing.Config) error {
	startValue, err := evaluate(fn, start, cfg.Epsilon)
	if err != nil {
		return err
	}

	run, err := experiment.RunAnnealing(fn, start, cfg)
	if err != nil {
		return err
	}

	fmt.Printf("\n--- %s ---\n", name)
	fmt.Printf("Стартовая точка:     %v\n", start)
	fmt.Printf("Стартовое значение:  %.8f\n", startValue)
	fmt.Printf("Лучшая точка:        %v\n", run.Result.BestPoint)
	fmt.Printf("Лучшее значение:     %.8f\n", run.Result.BestValue)
	fmt.Printf("Итераций:            %d\n", run.Result.Iterations)
	fmt.Printf("Принятых переходов:  %d\n", run.Result.AcceptedMoves)
	fmt.Printf("Финальная температура: %.10f\n", run.Result.FinalTemperature)
	fmt.Printf("Расстояние до [1..1]:  %.8f\n", distanceToOnes(run.Result.BestPoint))
	fmt.Printf("Время выполнения:    %s\n", run.Duration)

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
		diff := x - 1
		sum += diff * diff
	}

	return math.Sqrt(sum)
}
