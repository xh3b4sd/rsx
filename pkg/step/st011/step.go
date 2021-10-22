package st011

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

// mutate: <amount> RSX market cap
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Treasury.RSX.Supply.MarketCap = ctx.Treasury.RSX.Supply.Total * ctx.RSX.Price.Ceiling

	return ctx, nil
}
