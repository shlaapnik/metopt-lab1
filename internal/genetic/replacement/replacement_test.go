package replacement

import (
	"testing"

	"metopt-lab1/internal/genetic"
)

func TestAbsentReplacement(t *testing.T) {
	repl := &AbsentReplacement[float64]{}
	oldPop := &genetic.Population[float64]{
		Individuals: []genetic.Individual[float64]{{Score: 100}},
	}
	offspring := []genetic.Individual[float64]{{Score: 10}, {Score: 20}}

	result := repl.Replace(oldPop, offspring)

	if len(result) != len(offspring) {
		t.Errorf("Expected %d individuals, got %d", len(offspring), len(result))
	}
	if result[0].Score != offspring[0].Score {
		t.Error("AbsentReplacement modified offspring data")
	}
}

func TestElitismReplacement(t *testing.T) {
	repl := NewElitismReplacement[float64](2)
	oldPop := &genetic.Population[float64]{
		Individuals: []genetic.Individual[float64]{
			{Score: 10, Genes: []float64{1}},
			{Score: 50, Genes: []float64{2}},
			{Score: 100, Genes: []float64{3}},
		},
	}
	offspring := []genetic.Individual[float64]{{Score: 1}}

	result := repl.Replace(oldPop, offspring)

	if len(result) != 3 {
		t.Errorf("Expected 3 individuals, got %d", len(result))
	}

	if result[0].Score != 100 || result[1].Score != 50 {
		t.Errorf("Elitism failed to pick best individuals. Got scores: %v, %v", result[0].Score, result[1].Score)
	}
}
