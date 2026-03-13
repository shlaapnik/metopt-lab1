package replacement

import "metopt-lab1/internal/genetic"

type ElitismReplacement[V genetic.Number] struct {
	EliteCount int
}

func NewElitismReplacement[V genetic.Number](count int) *ElitismReplacement[V] {
	return &ElitismReplacement[V]{EliteCount: count}
}

func (e *ElitismReplacement[V]) Replace(oldPop *genetic.Population[V], offspring []genetic.Individual[V]) []genetic.Individual[V] {
	oldPop.Sort()

	count := e.EliteCount
	if count > len(oldPop.Individuals) {
		count = len(oldPop.Individuals)
	}

	result := make([]genetic.Individual[V], 0, count+len(offspring))

	for i := 0; i < count; i++ {
		elite := oldPop.Individuals[i]
		elite.Genes = genetic.CloneGenes(elite.Genes)
		result = append(result, elite)
	}

	result = append(result, offspring...)

	return result
}
