package genetic

import (
	"math/rand"
	"time"
)

func NewRand() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func CloneGenes(genes []float64) []float64 {
	d := make([]float64, len(genes))
	copy(d, genes)
	return d
}
