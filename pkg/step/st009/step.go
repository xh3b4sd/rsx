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

// verify: set <Value> RSX total supply
func (s Step) Run(ctx context.Context) (context.Context, error) {
	var amo float64
	{
		amo += ctx.Treasury.RSX.Amount

		amo = round.RoundN(amo, 4)
	}

	if amo != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, amo)
	}

	return ctx, nil
}
