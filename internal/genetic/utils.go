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

func CloneGenes[V comparable](genes []V) []V {
	d := make([]V, len(genes))
	copy(d, genes)
	return d
}

func Clamp[V Number](val, min, max V) V {
	if val < min {
		return min
	}
	if val > max {
		return max
	}
	return val
}
