package st010

import (
	"github.com/xh3b4sd/tracer"

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

// verify: set <Value> RSX market cap
func (s Step) Run(ctx context.Context) (context.Context, error) {
	if ctx.Pool.RSX.MarketCap != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, ctx.Pool.RSX.MarketCap)
	}

	return ctx, nil
}
