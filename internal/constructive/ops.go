package constructive

func Add(a, b Real) Real {
	return NewFromApprox(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrInvalidEpsilon
		}

		aApprox, err := a.Approx(eps / 2)
		if err != nil {
			return 0, err
		}

		bApprox, err := b.Approx(eps / 2)
		if err != nil {
			return 0, err
		}

		return aApprox + bApprox, nil
	})
}

func Sub(a, b Real) Real {
	return NewFromApprox(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrInvalidEpsilon
		}

		aApprox, err := a.Approx(eps / 2)
		if err != nil {
			return 0, err
		}

		bApprox, err := b.Approx(eps / 2)
		if err != nil {
			return 0, err
		}

		return aApprox - bApprox, nil
	})
}
func Mul(a, b Real) Real {
	return NewFromApprox(func(eps float64) (float64, error) {
		if eps <= 0 {
			return 0, ErrInvalidEpsilon
		}

		aApprox, err := a.Approx(eps / 4)
		if err != nil {
			return 0, err
		}

		bApprox, err := b.Approx(eps / 4)
		if err != nil {
			return 0, err
		}

		return aApprox * bApprox, nil
	})
}
