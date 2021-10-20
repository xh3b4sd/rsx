package st022

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

// mutate: rem <Value> DAI from protocol debt
func (s Step) Run(ctx context.Context) (context.Context, error) {
	if ctx.Protocol.RSX.Debt.Amount == 0 {
		return ctx, nil
	}

	// Every time the price floor increases the total value owed in debt
	// increases proportionally. That is why it is most important to pay back
	// debt as early as possible.
	ctx.Protocol.RSX.Debt.Value = ctx.Protocol.RSX.Debt.Amount * ctx.RSX.Price.Floor

	exc := ctx.Treasury.DAI.Excess * s.Value
	rsx := exc / ctx.RSX.Price.Floor
	rde := ctx.Protocol.RSX.Debt.Amount - rsx
	ede := ctx.Protocol.RSX.Debt.Value - exc

	ctx.Treasury.DAI.Excess -= exc
	ctx.Treasury.RSX.Supply.Circulating += rsx

	if rde < 0 || ede < 0 {
		rde *= -1
		ctx.Protocol.RSX.Debt.Amount = 0
		ctx.Protocol.RSX.Debt.Value = 0
		ctx.Treasury.DAI.Excess += (rde * ctx.RSX.Price.Floor)
	} else {
		ctx.Protocol.RSX.Debt.Amount -= rsx
		ctx.Protocol.RSX.Debt.Value -= exc
	}

	return ctx, nil
}
