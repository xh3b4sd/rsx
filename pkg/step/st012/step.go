package st012

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

// mutate: buy RSX for <Value> DAI from pool
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.Pool.RSXDAI.DAI.Amount += s.Value
	ctx.Pool.RSXDAI.DAI.Price = 1
	ctx.Pool.RSXDAI.DAI.Value += s.Value

	ctx.Pool.RSXDAI.RSX.Amount = ctx.Pool.RSXDAI.ConstantK / ctx.Pool.RSXDAI.DAI.Amount
	ctx.Pool.RSXDAI.RSX.Price = ctx.Pool.RSXDAI.DAI.Amount / ctx.Pool.RSXDAI.RSX.Amount
	ctx.Pool.RSXDAI.RSX.Value = ctx.Pool.RSXDAI.RSX.Amount / ctx.Pool.RSXDAI.RSX.Price

	return ctx, nil
}
