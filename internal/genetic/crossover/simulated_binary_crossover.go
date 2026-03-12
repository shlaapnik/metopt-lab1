package crossover

import (
	"math"
	"math/rand"

	"metopt-lab1/internal/genetic"
)

type SBXCrossover[V genetic.Number] struct {
	Prob   float64
	Eta    float64
	Bounds [][]float64
	rand   *rand.Rand
}

func NewSBXCrossover[V genetic.Number](prob float64, eta float64, bounds [][]float64, seed int) *SBXCrossover[V] {
	return &SBXCrossover[V]{
		Prob:   prob,
		Eta:    eta,
		Bounds: bounds,
		rand:   genetic.NewRand(seed),
	}
}

func (c *SBXCrossover[V]) Mate(p1, p2 genetic.Individual[V]) (genetic.Individual[V], genetic.Individual[V]) {
	if rand.Float64() > c.Prob {
		return genetic.Individual[V]{Genes: genetic.CloneGenes(p1.Genes)}, genetic.Individual[V]{Genes: genetic.CloneGenes(p2.Genes)}
	}

	dim := len(p1.Genes)
	child1 := genetic.Individual[V]{Genes: make([]V, dim)}
	child2 := genetic.Individual[V]{Genes: make([]V, dim)}

	for i := 0; i < dim; i++ {
		v1, v2 := float64(p1.Genes[i]), float64(p2.Genes[i])

		if math.Abs(v1-v2) < 1e-10 {
			child1.Genes[i], child2.Genes[i] = p1.Genes[i], p2.Genes[i]
			continue
		}

		u := rand.Float64()
		var beta float64
		if u <= 0.5 {
			beta = math.Pow(2.0*u, 1.0/(c.Eta+1.0))
		} else {
			beta = math.Pow(1.0/(2.0*(1.0-u)), 1.0/(c.Eta+1.0))
		}

		c1 := 0.5 * ((1.0+beta)*v1 + (1.0-beta)*v2)
		c2 := 0.5 * ((1.0-beta)*v1 + (1.0+beta)*v2)

		if i < len(c.Bounds) {
			c1 = genetic.Clamp(c1, c.Bounds[i][0], c.Bounds[i][1])
			c2 = genetic.Clamp(c2, c.Bounds[i][0], c.Bounds[i][1])
		}

		child1.Genes[i] = V(c1)
		child2.Genes[i] = V(c2)
	}

	return child1, child2
}
