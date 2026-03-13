package selection

import (
	"testing"

	"metopt-lab1/internal/genetic"
)

func TestTournamentSelection(t *testing.T) {
	sel := NewTournament[float64](100, 0)
	pop := &genetic.Population[float64]{
		Individuals: []genetic.Individual[float64]{
			{Score: 0.1},
			{Score: 0.2},
			{Score: 999.0},
		},
	}

	selected := sel.Select(pop, 5)

	if len(selected) != 5 {
		t.Errorf("Expected 5 selected individuals, got %d", len(selected))
	}
}

func TestRankingSelection(t *testing.T) {
	sel := NewRanking[float64](0)
	pop := &genetic.Population[float64]{
		Individuals: []genetic.Individual[float64]{
			{Score: 10.0},
			{Score: 20.0},
			{Score: 30.0},
		},
	}

	k := 50
	selected := sel.Select(pop, k)

	if len(selected) != k {
		t.Errorf("Expected %d selected individuals, got %d", k, len(selected))
	}

	if pop.Individuals[0].Score < pop.Individuals[1].Score {
		t.Error("Ranking selector did not sort the population")
	}
}
