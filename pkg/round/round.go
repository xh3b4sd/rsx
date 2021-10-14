package round

import (
	"fmt"
	"strconv"
)

func Round(f float64, p uint) float64 {
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", p)+"f", f)

	flo, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(err)
	}

	return flo
}
