package st005

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

// mutate: add <Value> RSX / DAI liquidity to pool
func (s Step) Run(ctx context.Context) (context.Context, error) {
	val := s.Value / 2

	ctx.Treasury.RSX.Minted += val / ctx.RSX.Price.Ceiling

	return ctx, nil
}
