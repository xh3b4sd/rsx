package st022

import (
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/xh3b4sd/rsx/pkg/context"
)

func Test_st022_1_percent(t *testing.T) {
	testCases := []struct {
		i func() context.Context
		o func() context.Context
	}{
		// Case 0
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 220
				ctx.Protocol.RSX.Debt.Value = 220
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 10000
				ctx.Treasury.RSX.Supply.Circulating = 220

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 120
				ctx.Protocol.RSX.Debt.Value = 120
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9900
				ctx.Treasury.RSX.Supply.Circulating = 320

				return ctx
			},
		},
		// Case 1
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 120
				ctx.Protocol.RSX.Debt.Value = 120
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9900
				ctx.Treasury.RSX.Supply.Circulating = 320

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 21
				ctx.Protocol.RSX.Debt.Value = 21
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9801
				ctx.Treasury.RSX.Supply.Circulating = 419

				return ctx
			},
		},
		// Case 2
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 21
				ctx.Protocol.RSX.Debt.Value = 21
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9801
				ctx.Treasury.RSX.Supply.Circulating = 419

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9780
				ctx.Treasury.RSX.Supply.Circulating = 517.01

				return ctx
			},
		},
		// Case 3
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9780
				ctx.Treasury.RSX.Supply.Circulating = 517.01

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 1.00
				ctx.Treasury.DAI.Excess = 9780
				ctx.Treasury.RSX.Supply.Circulating = 517.01

				return ctx
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var s Step
			{
				s = Step{
					Value: 0.01,
				}
			}

			o, err := s.Run(tc.i())
			if err != nil {
				t.Fatal(err)
			}

			if o != tc.o() {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.o(), o))
			}
		})
	}
}

func Test_st022_2_percent(t *testing.T) {
	testCases := []struct {
		i func() context.Context
		o func() context.Context
	}{
		// Case 0
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 733.3333333333
				ctx.Protocol.RSX.Debt.Value = 19.999999999989967
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 10000
				ctx.Treasury.RSX.Supply.Circulating = 733.3333333333

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 66.6666666666332
				ctx.Protocol.RSX.Debt.Value = 19.999999999989967
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 9800
				ctx.Treasury.RSX.Supply.Circulating = 1399.9999999999668

				return ctx
			},
		},
		// Case 1
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 66.6666666666332
				ctx.Protocol.RSX.Debt.Value = 19.999999999989967
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 9800
				ctx.Treasury.RSX.Supply.Circulating = 1399.9999999999668

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 9780.000000000011
				ctx.Treasury.RSX.Supply.Circulating = 2053.3333333333003

				return ctx
			},
		},
		// Case 2
		{
			i: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 9780.000000000011
				ctx.Treasury.RSX.Supply.Circulating = 2053.3333333333003

				return ctx
			},
			o: func() context.Context {
				ctx := context.Context{}

				ctx.Protocol.RSX.Debt.Amount = 0
				ctx.Protocol.RSX.Debt.Value = 0
				ctx.RSX.Price.Floor = 0.30
				ctx.Treasury.DAI.Excess = 9780.000000000011
				ctx.Treasury.RSX.Supply.Circulating = 2053.3333333333003

				return ctx
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			var s Step
			{
				s = Step{
					Value: 0.02,
				}
			}

			o, err := s.Run(tc.i())
			if err != nil {
				t.Fatal(err)
			}

			if o != tc.o() {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.o(), o))
			}
		})
	}
}
