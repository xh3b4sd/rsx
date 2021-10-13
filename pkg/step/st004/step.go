package st004

import "github.com/xh3b4sd/rsx/pkg/context"

type Step struct {
	Comment string
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

// add <Value> DAI to treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Protocol.Debt.RSX.Price = 2
	ctx.Protocol.Debt.RSX.Amount = s.Value / 2
	ctx.Protocol.Debt.RSX.Value = s.Value

	ctx.Treasury.DAI.Price = 1
	ctx.Treasury.DAI.Amount = s.Value / 1
	ctx.Treasury.DAI.Value = s.Value

	return ctx, nil
}
