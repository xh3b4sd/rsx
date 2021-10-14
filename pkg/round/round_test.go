package round

import (
	"strconv"
	"testing"
)

func Test_Round(t *testing.T) {
	testCases := []struct {
		f float64
		p uint
		r float64
	}{
		// Case 0
		{
			f: 3.14,
			p: 2,
			r: 3.14,
		},
		// Case 1
		{
			f: 3.44444,
			p: 3,
			r: 3.444,
		},
		// Case 2
		{
			f: 3.5,
			p: 0,
			r: 4,
		},
		// Case 3
		{
			f: 3.487,
			p: 8,
			r: 3.487,
		},
		// Case 4
		{
			f: 3.487,
			p: 2,
			r: 3.49,
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			r := Round(tc.f, tc.p)
			if r != tc.r {
				t.Fatal("expected", tc.r, "got", r)
			}
		})
	}
}
