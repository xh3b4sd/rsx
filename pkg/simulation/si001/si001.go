package si001

import (
	"fmt"

	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/step"
	"github.com/xh3b4sd/rsx/pkg/step/st002"
	"github.com/xh3b4sd/rsx/pkg/step/st003"
	"github.com/xh3b4sd/rsx/pkg/step/st004"
	"github.com/xh3b4sd/rsx/pkg/step/st005"
	"github.com/xh3b4sd/rsx/pkg/step/st006"
	"github.com/xh3b4sd/rsx/pkg/step/st007"
	"github.com/xh3b4sd/rsx/pkg/step/st008"
	"github.com/xh3b4sd/rsx/pkg/step/st009"
	"github.com/xh3b4sd/rsx/pkg/step/st010"
	"github.com/xh3b4sd/rsx/pkg/step/st011"
)

func Run() error {
	var err error

	ctx := context.Context{}

	steps := []step.Interface{
		st002.Step{Value: 1.00, Comment: "set price floor to 1 DAI"},
		st003.Step{Value: 2.00, Comment: "set price ceiling to 2 DAI"},
		st004.Step{Value: 2e06, Comment: "add 2M DAI to treasury"},
		st005.Step{Value: 4e06, Comment: "add 4M RSX / DAI liquidity to pool"},
		st006.Step{Value: 4e06, Comment: "ensure 4M seed investment"},
		st007.Step{Value: 4e06, Comment: "ensure 4M protocol debt"},
		st008.Step{Value: 1e06, Comment: "ensure 1M RSX circulating supply"},
		st009.Step{Value: 3e06, Comment: "ensure 3M RSX total supply"},
		st010.Step{Value: true, Comment: "ensure all circulating RSX is backed"},
		st011.Step{Value: 5e04, Comment: "buy RSX for 50k DAI from pool"},
	}

	for _, s := range steps {
		fmt.Println(s.Com())

		ctx, err = s.Run(ctx)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
