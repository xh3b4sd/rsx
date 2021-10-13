package st007

import (
	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/tracer"
)

type Step struct {
	Comment string
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

// ensure <Value> protocol debt
func (s Step) Run(ctx context.Context) (context.Context, error) {
	val := ctx.Protocol.Debt.RSX.Value

	if val != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, val)
	}

	return ctx, nil
}
