package selection

import (
	"math/rand"

	"github.com/shlaapnik/metopt-lab1/internal/genetic"
)

type Tournament[V genetic.Number] struct {
	Size int
	rand *rand.Rand
}

func NewTournament[V genetic.Number](size int) *Tournament[V] {
	if size <= 0 {
		size = 3
	}
	return &Tournament[V]{Size: size, rand: genetic.NewRand()}
}

func (t *Tournament[V]) Select(pop *genetic.Population[V], k int) []genetic.Individual[V] {
	selected := make([]genetic.Individual[V], 0, k)

	for i := 0; i < k; i++ {
		bestIdx := -1

		for j := 0; j < t.Size; j++ {
			idx := t.rand.Intn(len(pop.Individuals))

			if bestIdx == -1 || pop.Individuals[idx].Score > pop.Individuals[bestIdx].Score {
				bestIdx = idx
			}
		}

		winner := pop.Individuals[bestIdx]
		winner.Genes = genetic.CloneGenes(winner.Genes)
		selected = append(selected, winner)
	}

	return selected
}
