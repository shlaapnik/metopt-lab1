package genetic

import (
	"testing"
)

func TestPopulation_Best(t *testing.T) {
	pop := &Population[float64]{
		Individuals: []Individual[float64]{
			{Score: 10.0},
			{Score: 50.0},
			{Score: 5.0},
		},
	}

	best := pop.Best()
	if best.Score != 50.0 {
		t.Errorf("Best() gave Score = %f; want 50.0", best.Score)
	}
}

func TestPopulation_Sort(t *testing.T) {
	pop := &Population[float64]{
		Individuals: []Individual[float64]{
			{Score: 10.0},
			{Score: 50.0},
			{Score: 5.0},
			{Score: 100.0},
		},
	}

	pop.Sort()

	expected := []float64{100.0, 50.0, 10.0, 5.0}
	for i, indiv := range pop.Individuals {
		if indiv.Score != expected[i] {
			t.Errorf("At position %d expected Score %f, received %f", i, expected[i], indiv.Score)
		}
	}
}
