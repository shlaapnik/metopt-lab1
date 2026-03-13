package genetic

import (
	"testing"
)

func TestEngine_evaluate(t *testing.T) {
	fitnessFunc := func(genes []float64) float64 {
		return genes[0] + genes[1]
	}

	engine := &Engine[float64]{
		Fitness: fitnessFunc,
		Weight:  2.0,
	}

	inds := []Individual[float64]{
		{Genes: []float64{1.0, 2.0}},
		{Genes: []float64{3.0, 4.0}},
	}

	engine.evaluate(inds)

	if inds[0].Fitness != 3.0 || inds[0].Score != 6.0 {
		t.Errorf("Individual 0 assessment error: Fitness=%f, Score=%f", inds[0].Fitness, inds[0].Score)
	}

	if inds[1].Fitness != 7.0 || inds[1].Score != 14.0 {
		t.Errorf("Individual 1 assessment error: Fitness=%f, Score=%f", inds[1].Fitness, inds[1].Score)
	}
}

func TestEngine_initPopulation(t *testing.T) {
	bounds := [][]float64{{0.0, 10.0}, {-5.0, 5.0}}
	generator := func(index int, b []float64) float64 {
		return b[0]
	}

	engine := &Engine[float64]{
		PopSize:   2,
		Bounds:    bounds,
		Generator: generator,
	}

	engine.initPopulation()

	if len(engine.Population.Individuals) != 2 {
		t.Fatalf("The expected population size was 2, received %d", len(engine.Population.Individuals))
	}

	for _, indiv := range engine.Population.Individuals {
		if len(indiv.Genes) != 2 {
			t.Fatalf("The expected len of genes was 2, received %d", len(indiv.Genes))
		}
		if indiv.Genes[0] != 0.0 || indiv.Genes[1] != -5.0 {
			t.Errorf("Incorrect generation of genes: %v", indiv.Genes)
		}
	}
}
