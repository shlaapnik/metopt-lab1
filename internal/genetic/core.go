package genetic

import "sort"

type Individual struct {
	Genes   []float64
	Fitness float64
	Score   float64
}

type Population struct {
	Individuals []Individual
}

func (p *Population) Best() Individual {
	best := p.Individuals[0]
	for _, indiv := range p.Individuals {
		if indiv.Fitness > best.Fitness {
			best = indiv
		}
	}

	return best
}

func (p *Population) Sort() {
	sort.Slice(p.Individuals, func(i, j int) bool {
		return p.Individuals[i].Score > p.Individuals[j].Score
	})
}

type Selector interface {
	Select(popul *Population, k int) []Individual
}

type CrossoverType interface {
	Mate(p1, p2 Individual) (Individual, Individual)
}

type Mutator interface {
	Mutate(indiv *Individual, rate float64)
}

type FitnessFunc func(genes []float64) float64
