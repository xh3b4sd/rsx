package st005

import (
	"math"

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

// add <Value> RSX / DAI liquidity to pool
func (s Step) Run(ctx context.Context) (context.Context, error) {
	val := s.Value / 2

	ctx.Pool.RSXDAI.RSX.Amount = val / 2
	ctx.Pool.RSXDAI.RSX.Price = 2
	ctx.Pool.RSXDAI.RSX.Value = val

	ctx.Pool.RSXDAI.DAI.Amount = val / 1
	ctx.Pool.RSXDAI.DAI.Price = 1
	ctx.Pool.RSXDAI.DAI.Value = val

	ctx.Pool.RSXDAI.ConstantK = ctx.Pool.RSXDAI.RSX.Amount * ctx.Pool.RSXDAI.DAI.Amount
	ctx.Pool.RSXDAI.Liquidity = math.Sqrt(ctx.Pool.RSXDAI.RSX.Amount * ctx.Pool.RSXDAI.DAI.Amount)

	ctx.Protocol.Debt.RSX.Amount = s.Value / 2
	ctx.Protocol.Debt.RSX.Price = 2
	ctx.Protocol.Debt.RSX.Value = s.Value

	return ctx, nil
}
