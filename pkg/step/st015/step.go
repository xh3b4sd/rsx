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

// arb RSX for <Value> DAI between protocol and pool
func (s Step) Run(ctx context.Context) (context.Context, error) {
	amo := s.Value / ctx.RSX.Price.Ceiling

	{
		ctx.Treasury.ExcessReserves += amo
		ctx.Treasury.VolumeInflow += s.Value

		ctx.Treasury.DAI.Amount += s.Value
		ctx.Treasury.RSX.Amount += amo
	}

	{
		ctx.Pool.RSXDAI.RSX.Amount += amo
		ctx.Pool.RSXDAI.DAI.Amount = ctx.Pool.RSXDAI.ConstantK / ctx.Pool.RSXDAI.RSX.Amount

		ctx.Pool.RSXDAI.RSX.Price = ctx.Pool.RSXDAI.DAI.Amount / ctx.Pool.RSXDAI.RSX.Amount
		ctx.Pool.RSXDAI.DAI.Price = 1

		ctx.Pool.RSXDAI.RSX.Value = ctx.Pool.RSXDAI.RSX.Amount / ctx.Pool.RSXDAI.RSX.Price
		ctx.Pool.RSXDAI.DAI.Value = ctx.Pool.RSXDAI.DAI.Amount * ctx.Pool.RSXDAI.DAI.Price
	}

	return ctx, nil
}
