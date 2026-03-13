package selection

import (
	"math/rand"

	"metopt-lab1/internal/genetic"
)

type Ranking[V genetic.Number] struct {
	rand *rand.Rand
}

func NewRanking[V genetic.Number](seed int) *Ranking[V] {

	return &Ranking[V]{rand: genetic.NewRand(seed)}
}

func (r *Ranking[V]) Select(pop *genetic.Population[V], k int) []genetic.Individual[V] {
	pop.Sort()
	n := len(pop.Individuals)

	rankSum := float64(n * (n + 1) / 2)

	selected := make([]genetic.Individual[V], 0, k)

	for i := 0; i < k; i++ {
		pick := r.rand.Float64() * rankSum
		currentSum := 0.0

		for idx, ind := range pop.Individuals {
			rank := float64(n - idx)
			currentSum += rank

			if currentSum >= pick {
				chosen := ind
				chosen.Genes = genetic.CloneGenes(ind.Genes)
				selected = append(selected, chosen)
				break
			}
		}
	}

	return selected
}
