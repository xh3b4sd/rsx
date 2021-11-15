package bond

const (
	DISCOUNT_PERCENTAGE = 10
)

const (
	VOLUME_CAPACITY = 1000
)

// Discount returns the absolute bond price given the current price, floor and
// ceiling. The discount percentage is globally set and scales along the current
// price curve. The same discount is applied only limited by floor and ceiling.
func Discount(p float64, f float64, c float64) float64 {
	if p >= c {
		return c
	}

	if p <= f {
		return f
	}

	d := p * (100 - DISCOUNT_PERCENTAGE) / 100

	if d <= f {
		return f
	}

	return d
}

// Volume returns the absolute volume capacity for a bond given the current
// price, floor and ceiling. Volume capacity is globally set for a given time
// window, e.g. 24 hours. The volume capacity returned here scales along the
// current price curve. Lower prices allow lower volume capacity.
func Volume(p float64, f float64, c float64) float64 {
	if p >= c {
		return 0
	}

	if p <= f {
		return 0
	}

	return (p - f) / (c - f) * VOLUME_CAPACITY
}
