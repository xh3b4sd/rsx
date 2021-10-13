package st006

import (
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
)

type Step struct {
	Comment string
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

// ensure <Value> in seed investment
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var val float64
	{
		val += ctx.Pool.RSXDAI.DAI.Value
		val += ctx.Pool.RSXOHM.OHM.Value
		val += ctx.Treasury.DAI.Value
	}

	if val != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, val)
	}

	return ctx, nil
}
