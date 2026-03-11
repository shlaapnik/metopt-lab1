package genetic

import "math/rand"

type Engine[V Number] struct {
	PopSize     int
	Generations int

	Bounds [][]V
	Weight float64

	Fitness     FitnessFunc[V]
	Generator   GeneGenerator[V]
	Selector    Selector[V]
	Crossover   Crossover[V]
	Mutator     Mutator[V]
	Replacement Replacement[V]

	Population *Population[V]
	rand       *rand.Rand
}

func NewEngine[V Number](
	bounds [][]V,
	fitness FitnessFunc[V],
	geneGenerator GeneGenerator[V],
	selector Selector[V],
	crossover Crossover[V],
	mutator Mutator[V],
	replacement Replacement[V],
	popSize int,
	generations int,
	weight float64,
) *Engine[V] {
	if popSize <= 0 {
		popSize = 100
	}
	if generations <= 0 {
		generations = 75
	}
	if weight == 0 {
		weight = -1.0
	}

	return &Engine[V]{
		PopSize:     popSize,
		Generations: generations,
		Bounds:      bounds,
		Weight:      weight,
		Fitness:     fitness,
		Generator:   geneGenerator,
		Selector:    selector,
		Crossover:   crossover,
		Mutator:     mutator,
		Replacement: replacement,
		rand:        NewRand(),
	}
}

func (e *Engine[V]) initPopulation() {
	dim := len(e.Bounds)
	individuals := make([]Individual[V], e.PopSize)

	for i := 0; i < e.PopSize; i++ {
		genes := make([]V, dim)
		for j := 0; j < dim; j++ {
			genes[j] = e.Generator(j, e.Bounds[j])
		}
		individuals[i] = Individual[V]{Genes: genes}
	}
	e.Population = &Population[V]{Individuals: individuals}
}

func (e *Engine[V]) evaluate(inds []Individual[V]) {
	for i := range inds {
		inds[i].Fitness = e.Fitness(inds[i].Genes)
		inds[i].Score = inds[i].Fitness * e.Weight
	}
}

func (e *Engine[V]) Run() Individual[V] {
	e.initPopulation()
	e.evaluate(e.Population.Individuals)

	for gen := 0; gen < e.Generations; gen++ {
		parents := e.Selector.Select(e.Population, e.PopSize)
		offspring := make([]Individual[V], 0, e.PopSize)

		for i := 0; i < len(parents)-1; i += 2 {
			child1, child2 := e.Crossover.Mate(parents[i], parents[i+1])
			offspring = append(offspring, child1, child2)

		}

		if len(offspring) < e.PopSize {
			offspring = append(offspring, Individual[V]{Genes: CloneGenes(parents[len(parents)-1].Genes)})
		}

		for i := range offspring {
			e.Mutator.Mutate(&offspring[i])
		}

		e.evaluate(offspring)

		newIndividuals := e.Replacement.Replace(e.Population, offspring)

		if len(newIndividuals) > e.PopSize {
			newIndividuals = newIndividuals[:e.PopSize]
		}
		e.Population.Individuals = newIndividuals
	}

	return e.Population.Best()
}
