package experiment

import (
	"time"

	"github.com/shlaapnik/metopt-lab1/internal/annealing"
	"github.com/shlaapnik/metopt-lab1/internal/objective"
)

type AnnealingRun struct {
	Result   annealing.Result
	Duration time.Duration
}

func RunAnnealing(fn objective.Function, start []float64, cfg annealing.Config) (AnnealingRun, error) {
	startedAt := time.Now()
	result, err := annealing.Optimize(fn, start, cfg)
	if err != nil {
		return AnnealingRun{}, err
	}

	return AnnealingRun{
		Result:   result,
		Duration: time.Since(startedAt),
	}, nil
}
