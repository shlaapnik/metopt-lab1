package experiment

import (
	"time"

	"metopt-lab1/internal/genetic"
)

type GeneticRun[V genetic.Number] struct {
	Best     genetic.Individual[V]
	Duration time.Duration
}

func RunGenetic[V genetic.Number](e *genetic.Engine[V]) GeneticRun[V] {
	startedAt := time.Now()

	best := e.Run()

	return GeneticRun[V]{
		Best:     best,
		Duration: time.Since(startedAt),
	}
}
