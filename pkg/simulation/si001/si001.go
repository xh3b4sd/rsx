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
	"github.com/xh3b4sd/rsx/pkg/step/st008"
	"github.com/xh3b4sd/rsx/pkg/step/st009"
	"github.com/xh3b4sd/rsx/pkg/step/st010"
	"github.com/xh3b4sd/rsx/pkg/step/st011"
	"github.com/xh3b4sd/rsx/pkg/step/st012"
	"github.com/xh3b4sd/rsx/pkg/step/st013"
	"github.com/xh3b4sd/rsx/pkg/step/st014"
	"github.com/xh3b4sd/rsx/pkg/step/st015"
	"github.com/xh3b4sd/rsx/pkg/step/st016"
	"github.com/xh3b4sd/rsx/pkg/step/st017"
	"github.com/xh3b4sd/rsx/pkg/step/st018"
	"github.com/xh3b4sd/rsx/pkg/step/st019"
)

func Run() error {
	var err error

	ctx := context.Context{}

	steps := []step.Interface{
		st002.Step{Index: 0 /*****/, Value: 1.00 /*********/, Comment: "set price floor to 1.00 DAI"},
		st003.Step{Index: 1 /*****/, Value: 2.00 /*********/, Comment: "set price ceiling to 2.00 DAI"},

		st004.Step{Index: 2 /*****/, Value: 3e06 /*********/, Comment: "add 3M DAI to treasury"},
		st017.Step{Index: 3 /*****/, Value: 3e06 /*********/, Comment: "ensure 3M DAI in treasury"},
		st018.Step{Index: 4 /*****/, Value: 3e06 /*********/, Comment: "add 3M protocol debt in RSX"},
		st019.Step{Index: 5 /*****/, Value: 3e06 /*********/, Comment: "ensure 3M protocol debt in RSX"},

		st005.Step{Index: 6 /*****/, Value: 4e06 /*********/, Comment: "add 4M RSX / DAI liquidity to pool"},
		st008.Step{Index: 7 /*****/, Value: 1e06 /*********/, Comment: "ensure 1M RSX circulating supply"},
		st009.Step{Index: 8 /*****/, Value: 4e06 /*********/, Comment: "ensure 4M RSX total supply"},
		st010.Step{Index: 9 /*****/, Value: 2e06 /*********/, Comment: "ensure 2M RSX market cap"},
		st011.Step{Index: 10 /****/, Value: true /*********/, Comment: "ensure all circulating RSX is backed"},

		st013.Step{Index: 11 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},
		st014.Step{Index: 12 /****/, Value: 0 /************/, Comment: "ensure 0 excess reserves in treasury"},
		st016.Step{Index: 13 /****/, Value: 0 /************/, Comment: "ensure 0 volume inflow in treasury"},

		st012.Step{Index: 14 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 15 /****/, Value: 2.21 /*********/, Comment: "ensure RSX price of 2.21 DAI"},
		st015.Step{Index: 16 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 17 /****/, Value: 50e03 /********/, Comment: "ensure 50k excess reserves in treasury"},
		st013.Step{Index: 18 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 19 /****/, Value: 1e05 /*********/, Comment: "ensure 100k volume inflow in treasury"},

		st012.Step{Index: 20 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 21 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 22 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 23 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 24 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 25 /****/, Value: 10e04 /********/, Comment: "ensure 100k excess reserves in treasury"},
		st013.Step{Index: 26 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 27 /****/, Value: 2e05 /*********/, Comment: "ensure 200k volume inflow in treasury"},

		st012.Step{Index: 28 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 29 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 30 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 31 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 32 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 33 /****/, Value: 15e04 /********/, Comment: "ensure 150k excess reserves in treasury"},
		st013.Step{Index: 34 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 35 /****/, Value: 3e05 /*********/, Comment: "ensure 300k volume inflow in treasury"},

		st012.Step{Index: 36 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 37 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 38 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 39 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 40 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 41 /****/, Value: 20e04 /********/, Comment: "ensure 200k excess reserves in treasury"},
		st013.Step{Index: 42 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 43 /****/, Value: 4e05 /*********/, Comment: "ensure 400k volume inflow in treasury"},

		st012.Step{Index: 44 /****/, Value: 6000 /*********/, Comment: "buy RSX for 6k DAI from pool"},
		st013.Step{Index: 45 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 46 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 47 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 48 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 49 /****/, Value: 25e04 /********/, Comment: "ensure 250k excess reserves in treasury"},
		st013.Step{Index: 50 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 51 /****/, Value: 5e05 /*********/, Comment: "ensure 500k volume inflow in treasury"},

		st012.Step{Index: 52 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 53 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 54 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 55 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 56 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 57 /****/, Value: 30e04 /********/, Comment: "ensure 300k excess reserves in treasury"},
		st013.Step{Index: 58 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 59 /****/, Value: 6e05 /*********/, Comment: "ensure 600k volume inflow in treasury"},

		st012.Step{Index: 60 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 61 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 62 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 63 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 64 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 65 /****/, Value: 35e04 /********/, Comment: "ensure 350k excess reserves in treasury"},
		st013.Step{Index: 66 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 67 /****/, Value: 7e05 /*********/, Comment: "ensure 700k volume inflow in treasury"},

		st012.Step{Index: 68 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 69 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 70 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 71 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 72 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 73 /****/, Value: 40e04 /********/, Comment: "ensure 400k excess reserves in treasury"},
		st013.Step{Index: 74 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 75 /****/, Value: 8e05 /*********/, Comment: "ensure 800k volume inflow in treasury"},

		st012.Step{Index: 76 /****/, Value: 6000 /*********/, Comment: "buy RSX for 6k DAI from pool"},
		st013.Step{Index: 77 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 78 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 79 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 80 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 81 /****/, Value: 45e04 /********/, Comment: "ensure 450k excess reserves in treasury"},
		st013.Step{Index: 82 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 83 /****/, Value: 9e05 /*********/, Comment: "ensure 900k volume inflow in treasury"},

		st012.Step{Index: 84 /****/, Value: 4000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 85 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st012.Step{Index: 86 /****/, Value: 1e05 /*********/, Comment: "buy RSX for 100k DAI from pool"},
		st013.Step{Index: 87 /****/, Value: 2.20 /*********/, Comment: "ensure RSX price of 2.20 DAI"},
		st015.Step{Index: 88 /****/, Value: 1e05 /*********/, Comment: "arb RSX for 100k DAI between protocol and pool"},
		st014.Step{Index: 89 /****/, Value: 50e04 /********/, Comment: "ensure 500k excess reserves in treasury"},
		st013.Step{Index: 90 /****/, Value: 1.99 /*********/, Comment: "ensure RSX price of 1.99 DAI"},
		st016.Step{Index: 91 /****/, Value: 1e06 /*********/, Comment: "ensure 1M volume inflow in treasury"},

		st012.Step{Index: 92 /****/, Value: 6000 /*********/, Comment: "buy RSX for 4k DAI from pool"},
		st013.Step{Index: 93 /****/, Value: 2.00 /*********/, Comment: "ensure RSX price of 2.00 DAI"},

		st017.Step{Index: 94 /*****/, Value: 2e06 /********/, Comment: "ensure 2M DAI in treasury"},
		st019.Step{Index: 95 /*****/, Value: 3e06 /********/, Comment: "ensure 3M protocol debt in RSX"},
		st008.Step{Index: 96 /*****/, Value: 15e05 /*******/, Comment: "ensure 1.5M RSX circulating supply"},
		st009.Step{Index: 97 /*****/, Value: 45e05 /*******/, Comment: "ensure 4.5M RSX total supply"},
		st010.Step{Index: 98 /*****/, Value: 3e06 /********/, Comment: "ensure 3M RSX market cap"},
		st011.Step{Index: 99 /*****/, Value: true /********/, Comment: "ensure all circulating RSX is backed"},

		st002.Step{Index: 100 /****/, Value: 1.30 /********/, Comment: "set price floor to 1.30 DAI"},
		st003.Step{Index: 101 /****/, Value: 2.60 /********/, Comment: "set price ceiling to 2.60 DAI"},
		st014.Step{Index: 102 /****/, Value: 50e03 /*******/, Comment: "ensure 50k excess reserves in treasury"},
	}

	for i, s := range steps {
		if s.Ind() != i {
			return tracer.Maskf(executionFailedError, "expected %d, got %d", i, s.Ind())
		}

		var spc string
		{
			if i >= 0 && i <= 9 {
				spc = "      "
			}
			if i >= 10 && i <= 99 {
				spc = "     "
			}
			if i >= 100 && i <= 999 {
				spc = "    "
			}
		}

		fmt.Printf("%d%s%s\n", s.Ind(), spc, s.Com())

		ctx, err = s.Run(ctx)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
