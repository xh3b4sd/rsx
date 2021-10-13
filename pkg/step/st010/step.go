package st010

import (
	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/tracer"
)

type Step struct {
	Comment string
	Value   bool
}

func (s Step) Com() string {
	return s.Comment
}

// ensure all circulating RSX is backed
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var pri float64
	{
		var c float64

		if ctx.Pool.RSXDAI.RSX.Price != 0 {
			c++
			pri += ctx.Pool.RSXDAI.RSX.Price
		}

		if ctx.Pool.RSXOHM.RSX.Price != 0 {
			c++
			pri += ctx.Pool.RSXOHM.RSX.Price
		}

		pri /= c
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

	if val != ctx.Treasury.DAI.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", ctx.Treasury.DAI.Value, val)
	}

	return ctx, nil
}
