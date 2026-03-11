package genetic

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Individual[V Number] struct {
	Genes   []V
	Fitness float64
	Score   float64
}

type Population[V Number] struct {
	Individuals []Individual[V]
}

func (p *Population[V]) Best() Individual[V] {
	best := p.Individuals[0]
	for _, indiv := range p.Individuals {
		if indiv.Score > best.Score {
			best = indiv
		}
	}

	return best
}

func (p *Population[V]) Sort() {
	sort.Slice(p.Individuals, func(i, j int) bool {
		return p.Individuals[i].Score > p.Individuals[j].Score
	})
}

type Selector[V Number] interface {
	Select(pop *Population[V], k int) []Individual[V]
}

type Crossover[V Number] interface {
	Mate(p1, p2 Individual[V]) (Individual[V], Individual[V])
}

type Mutator[V Number] interface {
	Mutate(indiv *Individual[V])
}

type Replacement[V Number] interface {
	Replace(oldPop *Population[V], offspring []Individual[V]) []Individual[V]
}

type FitnessFunc[V Number] func(genes []V) float64

type GeneGenerator[V Number] func(index int, bounds []V) V
