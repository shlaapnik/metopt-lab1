package crossover

import (
	"math"
	"math/rand"

	"metopt-lab1/internal/genetic"
)

type BLXCrossover[V genetic.Number] struct {
	Prob   float64
	Alpha  float64
	Bounds [][]float64
	rand   *rand.Rand
}

func NewBlendCrossover[V genetic.Number](prob float64, alpha float64, bounds [][]float64, seed int) *BLXCrossover[V] {
	return &BLXCrossover[V]{
		Prob:   prob,
		Alpha:  alpha,
		Bounds: bounds,
		rand:   genetic.NewRand(seed),
	}
}

func (c *BLXCrossover[V]) Mate(p1, p2 genetic.Individual[V]) (genetic.Individual[V], genetic.Individual[V]) {
	if rand.Float64() > c.Prob {
		return genetic.Individual[V]{Genes: genetic.CloneGenes(p1.Genes)}, genetic.Individual[V]{Genes: genetic.CloneGenes(p2.Genes)}
	}

	dim := len(p1.Genes)
	child1 := genetic.Individual[V]{Genes: make([]V, dim)}
	child2 := genetic.Individual[V]{Genes: make([]V, dim)}

	for i := 0; i < dim; i++ {
		v1 := float64(p1.Genes[i])
		v2 := float64(p2.Genes[i])

		minP := math.Min(v1, v2)
		maxP := math.Max(v1, v2)
		diff := maxP - minP

		minRange := minP - c.Alpha*diff
		maxRange := maxP + c.Alpha*diff

		val1 := genetic.RandomFloat(minRange, maxRange)
		val2 := genetic.RandomFloat(minRange, maxRange)

		if i < len(c.Bounds) {
			val1 = genetic.Clamp(val1, c.Bounds[i][0], c.Bounds[i][1])
			val2 = genetic.Clamp(val2, c.Bounds[i][0], c.Bounds[i][1])
		}

		child1.Genes[i] = V(val1)
		child2.Genes[i] = V(val2)
	}

	return child1, child2
}
