package st018

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

// mutate: add <Value> protocol debt in RSX
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Protocol.Debt.RSX.Amount = s.Value
	ctx.Treasury.RSX.Amount = s.Value

	return ctx, nil
}
