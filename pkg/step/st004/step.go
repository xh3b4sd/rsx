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

// add <Value> DAI to treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Protocol.Debt.RSX.Amount = s.Value / ctx.RSX.Price.Floor
	ctx.Protocol.Debt.RSX.Price = ctx.RSX.Price.Floor
	ctx.Protocol.Debt.RSX.Value = s.Value

	ctx.Treasury.RSX.Amount += ctx.Protocol.Debt.RSX.Amount
	ctx.Treasury.DAI.Amount += ctx.Protocol.Debt.RSX.Value

	return ctx, nil
}
