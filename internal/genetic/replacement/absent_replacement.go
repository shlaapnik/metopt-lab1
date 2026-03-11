package replacement

import "github.com/shlaapnik/metopt-lab1/internal/genetic"

type AbsentReplacement[V genetic.Number] struct{}

func (s *AbsentReplacement[V]) Replace(oldPop *genetic.Population[V], offspring []genetic.Individual[V]) []genetic.Individual[V] {
	return offspring
}
