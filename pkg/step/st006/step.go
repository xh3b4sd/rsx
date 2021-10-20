package st006

import (
	"github.com/xh3b4sd/rsx/pkg/context"
)

type Step struct {
	Comment string
	Index   uint
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

func (s Step) Ind() int {
	return int(s.Index)
}

// mutate: <amount> RSX circulating supply
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var amo float64
	{
		amo += ctx.Treasury.RSX.Minted

		amo -= ctx.Protocol.RSX.Debt.Amount
	}

	ctx.Treasury.RSX.Supply.Circulating = amo

	return ctx, nil
}
