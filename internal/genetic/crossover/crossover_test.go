package crossover

import (
	"reflect"
	"testing"

	"metopt-lab1/internal/genetic"
)

func TestBLXCrossover_ProbZero(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}, {0.0, 10.0}}
	cross := NewBlendCrossover[float64](0.0, 0.5, bounds, 0)

	p1 := genetic.Individual[float64]{Genes: []float64{2.0, 3.0}}
	p2 := genetic.Individual[float64]{Genes: []float64{8.0, 7.0}}

	c1, c2 := cross.Mate(p1, p2)

	if !reflect.DeepEqual(p1.Genes, c1.Genes) {
		t.Errorf("In Prob=0.0 child 1 must be the copy of parent 1. Got: %v", c1.Genes)
	}
	if !reflect.DeepEqual(p2.Genes, c2.Genes) {
		t.Errorf("In Prob=0.0 child 2 must be the copy of parent 2. Got: %v", c2.Genes)
	}
}

func TestBLXCrossover_BoundsConstraint(t *testing.T) {
	bounds := [][]float64{{4.0, 6.0}, {4.0, 6.0}}
	cross := NewBlendCrossover[float64](1.0, 2.0, bounds, 0)

	p1 := genetic.Individual[float64]{Genes: []float64{4.5, 4.5}}
	p2 := genetic.Individual[float64]{Genes: []float64{5.5, 5.5}}

	c1, c2 := cross.Mate(p1, p2)

	for i := 0; i < 2; i++ {
		if c1.Genes[i] < bounds[i][0] || c1.Genes[i] > bounds[i][1] {
			t.Errorf("Child 1 gen out of bounds: %v (Bounds: %v)", c1.Genes, bounds)
		}
		if c2.Genes[i] < bounds[i][0] || c2.Genes[i] > bounds[i][1] {
			t.Errorf("Child 2 gen out of bounds: %v (Bounds: %v)", c2.Genes, bounds)
		}
	}
}

func TestSBXCrossover_BoundsConstraint(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}}
	cross := NewSBXCrossover[float64](1.0, 0.5, bounds, 0)

	p1 := genetic.Individual[float64]{Genes: []float64{1.0}}
	p2 := genetic.Individual[float64]{Genes: []float64{9.0}}

	c1, c2 := cross.Mate(p1, p2)

	if c1.Genes[0] < bounds[0][0] || c1.Genes[0] > bounds[0][1] {
		t.Errorf("Child 1 out of bounds: %v", c1.Genes)
	}
	if c2.Genes[0] < bounds[0][0] || c2.Genes[0] > bounds[0][1] {
		t.Errorf("Child 2 out of bounds: %v", c2.Genes)
	}
}
