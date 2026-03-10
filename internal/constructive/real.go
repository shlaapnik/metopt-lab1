package constructive

type Real interface {
	Approx(eps float64) (float64, error)
}

type realFunc struct {
	approx func(eps float64) (float64, error)
}

func NewFromFunc(f func(eps float64) (float64, error)) Real {
	return realFunc{approx: f}
}

func NewConstant(value float64) Real {
	return realFunc{
		approx: func(eps float64) (float64, error) {
			if eps <= 0 {
				return 0, ErrNonPositiveEpsilon
			}
			return value, nil
		},
	}
}

func (r realFunc) Approx(eps float64) (float64, error) {
	if eps <= 0 {
		return 0, ErrNonPositiveEpsilon
	}
	return r.approx(eps)
}
