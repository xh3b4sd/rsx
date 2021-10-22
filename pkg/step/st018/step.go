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
	// The initial protocol debt is incured for the seed investment of
	// bootstrapping the protocol. The debt is accounted for in RSX at a 50%
	// discount at price ceiling.
	ctx.Protocol.RSX.Debt.Amount = s.Value / ctx.RSX.Price.Ceiling * 1.5
	ctx.Protocol.RSX.Debt.Value = ctx.Protocol.RSX.Debt.Amount * ctx.RSX.Price.Floor
	ctx.Treasury.RSX.Minted = ctx.Protocol.RSX.Debt.Amount

	return ctx, nil
}
