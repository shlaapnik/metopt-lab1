package genetic

import "math/rand"

type Engine struct {
	PopSize     int
	Generations int

	Bounds [][]float64
	Weight float64

	Fitness     FitnessFunc
	Selector    Selector
	Crossover   Crossover
	Mutator     Mutator
	Replacement Replacement

	Population *Population
	rand       *rand.Rand
}

func NewEngine(
	bounds [][]float64,
	fitness FitnessFunc,
	selector Selector,
	crossover Crossover,
	mutator Mutator,
	replacement Replacement,
	popSize int,
	generations int,
	weight float64,
) *Engine {
	if popSize <= 0 {
		popSize = 100
	}
	if generations <= 0 {
		generations = 75
	}
	if weight == 0 {
		weight = -1.0
	}

	return &Engine{
		PopSize:     popSize,
		Generations: generations,
		Bounds:      bounds,
		Weight:      weight,
		Fitness:     fitness,
		Selector:    selector,
		Crossover:   crossover,
		Mutator:     mutator,
		Replacement: replacement,
		rand:        NewRand(),
	}
}

func (e *Engine) initPopulation() {
	dim := len(e.Bounds)
	individuals := make([]Individual, e.PopSize)

	for i := 0; i < e.PopSize; i++ {
		genes := make([]float64, dim)
		for j := 0; j < dim; j++ {
			genes[j] = RandomFloat(e.Bounds[j][0], e.Bounds[j][1])
		}
		individuals[i] = Individual{Genes: genes}
	}
	e.Population = &Population{Individuals: individuals}
}

func (e *Engine) evaluate(inds []Individual) {
	for i := range inds {
		inds[i].Fitness = e.Fitness(inds[i].Genes)
		inds[i].Score = inds[i].Fitness * e.Weight
	}
}

func (e *Engine) Run() Individual {
	e.initPopulation()
	e.evaluate(e.Population.Individuals)

	for gen := 0; gen < e.Generations; gen++ {
		parents := e.Selector.Select(e.Population, e.PopSize)
		offspring := make([]Individual, 0, e.PopSize)

		for i := 0; i < len(parents)-1; i += 2 {
			child1, child2 := e.Crossover.Mate(parents[i], parents[i+1])
			offspring = append(offspring, child1, child2)

		}

		if len(offspring) < e.PopSize {
			offspring = append(offspring, Individual{Genes: CloneGenes(parents[len(parents)-1].Genes)})
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
