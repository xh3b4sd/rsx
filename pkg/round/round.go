package round

import (
	"fmt"
	"math"
	"strconv"
)

func RoundP(f float64, p uint) float64 {
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", p)+"f", f)

	flo, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}

	return flo
}

func RoundN(f float64, n uint) float64 {
	p := math.Pow10(int(n))
	return RoundP(f/p, 0) * p
}
