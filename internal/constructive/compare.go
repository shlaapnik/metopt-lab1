package constructive

import "math"

func Compare(a, b Real, eps float64) (int, error) {
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

	diff := aApprox - bApprox
	if math.Abs(diff) <= eps {
		return 0, nil
	}
	if diff < 0 {
		return -1, nil
	}
	return 1, nil
}
