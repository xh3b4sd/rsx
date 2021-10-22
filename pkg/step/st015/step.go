package st015

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

// mutate: add <Value> DAI treasury inflow
func (s Step) Run(ctx context.Context) (context.Context, error) {
	amo := s.Value / ctx.RSX.Price.Ceiling

	{
		ctx.Treasury.DAI.Inflow += s.Value
		ctx.Treasury.DAI.Backing += s.Value
		ctx.Treasury.RSX.Minted += amo
	}

	return ctx, nil
}
