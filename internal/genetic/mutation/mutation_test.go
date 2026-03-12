package mutation

import (
	"reflect"
	"testing"

	"github.com/shlaapnik/metopt-lab1/internal/genetic"
)

func TestGaussianMutator_ProbZero(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}}
	mut := NewGaussianMutator[float64](0.0, bounds, 0.1)

	indiv := genetic.Individual[float64]{Genes: []float64{5.0}}
	original := []float64{5.0}

	mut.Mutate(&indiv)

	if !reflect.DeepEqual(indiv.Genes, original) {
		t.Errorf("In Prob=0.0 gen must not mutate. Got: %v", indiv.Genes)
	}
}

func TestGaussianMutator_EmptyGenes(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}}
	mut := NewGaussianMutator[float64](1.0, bounds, 0.1)

	indiv := genetic.Individual[float64]{Genes: []float64{}}

	mut.Mutate(&indiv)

	if len(indiv.Genes) != 0 {
		t.Errorf("Expected empty sclice of genes, got %d", len(indiv.Genes))
	}
}

func TestGaussianMutator_BoundsConstraint(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}}
	mut := NewGaussianMutator[float64](1.0, bounds, 100.0)

	for i := 0; i < 50; i++ {
		indiv := genetic.Individual[float64]{Genes: []float64{5.0}}
		mut.Mutate(&indiv)

		val := indiv.Genes[0]
		if val < bounds[0][0] || val > bounds[0][1] {
			t.Fatalf("Mutation took the gene out of bounds: %v (Iteration %d)", val, i)
		}
	}
}
