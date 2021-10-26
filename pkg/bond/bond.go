package bond

const (
	DISCOUNT_PERCENTAGE_MIN = 2
	DISCOUNT_PERCENTAGE_MAX = 5
)

func Discount(p float64, f float64, c float64) float64 {
	if p >= c {
		return 0
	}

	if p <= f {
		return 0
	}

	r := p * (100 - DISCOUNT_PERCENTAGE_MAX) / 100

	if r <= DISCOUNT_PERCENTAGE_MIN {
		if p*(100-DISCOUNT_PERCENTAGE_MIN)/100 <= f {
			return 0
		}

		return DISCOUNT_PERCENTAGE_MIN
	}

	return r
}

// TODO plot volume to see if it is linear
func Volume(p float64, f float64, c float64) float64 {
	if p >= c {
		return 0
	}

	if p <= f {
		return 0
	}

	return (p - f) / (c - f) * 100
}
