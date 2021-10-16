package st010

import (
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/round"
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

// verify: set <Value> RSX market cap
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var pri float64
	{
		pri = ctx.Pool.RSXPrice()
	}

	var amo float64
	{
		amo += ctx.Treasury.RSX.Minted

		amo -= ctx.Protocol.Debt.RSX.Amount
	}

	var val float64
	{
		val = amo * pri

		val = round.RoundN(val, 4)
	}

	if val != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, val)
	}

	return ctx, nil
}
