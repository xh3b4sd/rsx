package st021

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

// mutate: add <Value> DAI to DAO
func (s Step) Run(ctx context.Context) (context.Context, error) {
	val := ctx.Treasury.DAI.Excess * s.Value

	ctx.Treasury.DAI.Excess -= val
	ctx.Treasury.DAI.DAO += val

	return ctx, nil
}
