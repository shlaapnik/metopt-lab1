package objective

import "metopt-lab1/internal/constructive"

type Function interface {
	Dimension() int
	Evaluate(point []float64) (constructive.Real, error)
}
