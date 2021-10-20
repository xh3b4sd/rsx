package st020

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

// mutate: <amount> excess reserves in treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	// Calculate the desired amount of DAI backing RSX in circulation.
	var des float64
	{
		des = (ctx.Treasury.RSX.Supply.Circulating * ctx.RSX.Price.Floor)
	}

	// Calculate the remaining excess reserves when accounted for with required
	// RSX backing.
	var exc float64
	{
		exc = ctx.Treasury.DAI.Backing - des
	}

	ctx.Treasury.DAI.Excess = exc

	return ctx, nil
}
