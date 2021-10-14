package st009

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

// ensure <Value> RSX total supply
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var val float64
	{
		val += ctx.Pool.RSXDAI.RSX.Amount
		val += ctx.Pool.RSXOHM.RSX.Amount
		val += ctx.Protocol.Debt.RSX.Amount

		val = round.Round(val, 2)
	}

	if val != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, val)
	}

	return ctx, nil
}
