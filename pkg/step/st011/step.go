package st011

import (
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
)

type Step struct {
	Comment string
	Index   uint
	Value   bool
}

func (s Step) Com() string {
	return s.Comment
}

func (s Step) Ind() int {
	return int(s.Index)
}

// ensure all circulating RSX is backed
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var pri float64
	{
		pri = ctx.RSX.Price.Floor
	}

	var amo float64
	{
		amo += ctx.Pool.RSXDAI.RSX.Amount
		amo += ctx.Pool.RSXOHM.RSX.Amount
	}

	var val float64
	{
		val = amo * pri
	}

	if (ctx.Treasury.DAI.Amount >= val) != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", ctx.Treasury.DAI.Amount, val)
	}

	return ctx, nil
}
