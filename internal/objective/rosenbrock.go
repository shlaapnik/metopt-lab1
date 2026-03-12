package objective

import "metopt-lab1/internal/constructive"

type rosenbrock struct {
	dimension int
}

func NewRosenbrock2D() Function {
	return rosenbrock{dimension: 2}
}

func NewRosenbrock(dimension int) (Function, error) {
	if dimension < 2 {
		return nil, ErrInvalidDimension
	}
	return rosenbrock{dimension: dimension}, nil
}

func (r rosenbrock) Dimension() int {
	return r.dimension
}

func (r rosenbrock) Evaluate(point []float64) (constructive.Real, error) {
	if len(point) != r.dimension {
		return nil, ErrPointDimensionMismatch
	}

	value := 0.0
	for i := 0; i < r.dimension-1; i++ {
		xi := point[i]
		next := point[i+1]

		a := 1 - xi
		b := next - xi*xi
		value += a*a + 100*b*b
	}

	return constructive.NewConstant(value), nil
}
