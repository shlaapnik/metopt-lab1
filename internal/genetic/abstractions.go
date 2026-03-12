package genetic

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
