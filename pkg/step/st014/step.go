package st014

import (
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/round"
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

// ensure <Value> excess reserves in treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	// Calculate the RSX circulating supply. Note that this subtracts the
	// protocol debt as this is meant to be locked until accounted for by future
	// revenue.
	var rsa float64
	{
		rsa += ctx.Treasury.RSX.Amount

		rsa -= ctx.Protocol.Debt.RSX.Amount
	}

	// Calculate the amount of DAI in the treasury used for backing RSX in
	// circulation.
	var daa float64
	{
		daa = ctx.Treasury.DAI.Amount
	}

	// Calculate the desired amount of DAI backing RSX in circulation.
	var deb float64
	{
		deb = (rsa * ctx.RSX.Price.Floor)

		deb = round.RoundP(deb, 0)
	}

	// Calculate the remaining excess reserves when accounted for with required
	// RSX backing.
	var exc float64
	{
		exc = ctx.Treasury.DAI.Amount - deb

		ctx.Treasury.ExcessReserves = exc

		exc = round.RoundN(exc, 4)
	}

	// Calculate the current amount of DAI backing RSX in circulation.
	var cub float64
	{
		cub = daa - ctx.Treasury.ExcessReserves

		cub = round.RoundP(cub, 0)
	}

	if exc != s.Value {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", s.Value, exc)
	}
	if cub != deb {
		return context.Context{}, tracer.Maskf(executionFailedError, "expected %f, got %f", deb, cub)
	}

	return ctx, nil
}
