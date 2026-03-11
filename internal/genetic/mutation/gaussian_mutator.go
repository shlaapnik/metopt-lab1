package mutation

import (
	"math/rand"

	"github.com/shlaapnik/metopt-lab1/internal/genetic"
)

type GaussianMutator[V genetic.Number] struct {
	Prob   float64
	Bounds [][]float64
	Scale  float64
}

func NewGaussianMutator[V genetic.Number](prob float64, bounds [][]float64, scale float64) *GaussianMutator[V] {
	return &GaussianMutator[V]{
		Prob:   prob,
		Bounds: bounds,
		Scale:  scale,
	}
}

func (m *GaussianMutator[V]) Mutate(indiv *genetic.Individual[V]) {
	if rand.Float64() > m.Prob {
		return
	}

	dim := len(indiv.Genes)
	if dim == 0 {
		return
	}

	idx := rand.Intn(dim)

	minBound := m.Bounds[idx][0]
	maxBound := m.Bounds[idx][1]

	stdDev := (maxBound - minBound) * m.Scale

	newValue := rand.NormFloat64()*stdDev + float64(indiv.Genes[idx])

	clamped := genetic.Clamp(newValue, minBound, maxBound)
	indiv.Genes[idx] = V(clamped)
}
