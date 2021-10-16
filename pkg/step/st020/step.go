package st020

import (
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

// mutate: <amount> excess reserves in treasury
func (s Step) Run(ctx context.Context) (context.Context, error) {
	// Calculate the RSX circulating supply. Note that this subtracts the
	// protocol debt as this is meant to be locked until accounted for by future
	// revenue.
	var rsa float64
	{
		rsa += ctx.Treasury.RSX.Minted

		rsa -= ctx.Protocol.Debt.RSX.Amount
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
		exc = ctx.Treasury.DAI.Backing - deb
	}

	ctx.Treasury.DAI.Excess = exc

	return ctx, nil
}
