package st004

import "github.com/xh3b4sd/rsx/pkg/context"

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

// mutate: add <Value> DAI to treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Treasury.DAI.Amount += s.Value

	return ctx, nil
}
