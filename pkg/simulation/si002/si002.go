package si002

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/chart"
	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/step"
	"github.com/xh3b4sd/rsx/pkg/step/st002"
	"github.com/xh3b4sd/rsx/pkg/step/st003"
	"github.com/xh3b4sd/rsx/pkg/step/st004"
	"github.com/xh3b4sd/rsx/pkg/step/st005"
	"github.com/xh3b4sd/rsx/pkg/step/st006"
	"github.com/xh3b4sd/rsx/pkg/step/st007"
	"github.com/xh3b4sd/rsx/pkg/step/st011"
	"github.com/xh3b4sd/rsx/pkg/step/st012"
	"github.com/xh3b4sd/rsx/pkg/step/st015"
	"github.com/xh3b4sd/rsx/pkg/step/st018"
	"github.com/xh3b4sd/rsx/pkg/step/st020"
)

const (
	htmlFile = ".render/html/si002.html"
)

func Run() error {
	var err error

	{
		p := filepath.Dir(htmlFile)

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var chs []*charts.Line
	{
		chs, err = generate(context.Context{})
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := components.NewPage()

		for _, c := range chs {
			p.AddCharts(c)
		}

		f, err := os.Create(htmlFile)
		if err != nil {
			return tracer.Mask(err)
		}

		p.SetLayout(components.PageFlexLayout)

		err = p.Render(io.MultiWriter(f))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		log.Println("running server at http://localhost:8000")

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, htmlFile)
		})

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}

	return nil
}

func generate(ctx context.Context) ([]*charts.Line, error) {
	var floor float64
	{
		floor = 1.00
	}

	// multiple is the multiple to be applied to floor in order to compute the
	// actual price floor. E.g. 2.00 means price ceiling is always 2x price
	// floor.
	var ceiling float64
	var multiple float64
	{
		multiple = 2.00
	}

	// growth is the multiple at which price floor and price ceiling grow. E.g.
	// 0.10 means the price floor rises by 10% on each increase.
	var growth float64
	{
		growth = 0.02
	}

	var chs []*charts.Line

	var backing *chart.Chart
	{
		backing = chart.New("DAI Backing / Market Cap", "DAI Backing", "Market Cap")
		backing.MaxY(200e06)
		backing.MinY(0)
	}

	var excess *chart.Chart
	{
		excess = chart.New("Excess Reserves")
		excess.MaxY(15e06)
		excess.MinY(0)
	}

	var price *chart.Chart
	{
		price = chart.New("RSX Price", "Price Floor", "Price Ceiling")
		price.MaxY(20)
		price.MinY(0)
	}

	var supply *chart.Chart
	{
		supply = chart.New("RSX Supply", "Circulating Supply", "Total Supply")
		supply.MaxY(20e06)
		supply.MinY(0)
	}

	ctx = execute(ctx, []step.Interface{
		st002.Step{Value: 1.00 /***/, Comment: "mutate: set 1.00 DAI price floor"},
		st003.Step{Value: 2.00 /***/, Comment: "mutate: set 2.00 DAI price ceiling"},

		st004.Step{Value: 3e06 /***/, Comment: "mutate: add 3.0M DAI to treasury"},
		st018.Step{Value: 3e06 /***/, Comment: "mutate: add 3.0M protocol debt in RSX"},

		st005.Step{Value: 4e06 /***/, Comment: "mutate: add 4.0M RSX / DAI liquidity to pool"},
		st006.Step{ /***************/ Comment: "mutate: <amount> RSX circulating supply"},
		st007.Step{ /***************/ Comment: "mutate: <amount> RSX total supply"},
		st011.Step{ /***************/ Comment: "mutate: <amount> RSX market cap"},
	})

	for i := 0; i < 1000; i++ {
		{
			if backingCanIncrease(ctx, floor, growth) {
				floor = floor * (1 + growth)
				ceiling = floor * multiple

				ctx = execute(ctx, []step.Interface{
					st002.Step{Value: floor /*****/, Comment: fmt.Sprintf("mutate: set %.2f DAI price floor", floor)},
					st003.Step{Value: ceiling /***/, Comment: fmt.Sprintf("mutate: set %.2f DAI price ceiling", ceiling)},
					st020.Step{ /******************/ Comment: "mutate: <amount> excess reserves in treasury"},
					st006.Step{ /******************/ Comment: "mutate: <amount> RSX circulating supply"},
					st007.Step{ /******************/ Comment: "mutate: <amount> RSX total supply"},
					st011.Step{ /******************/ Comment: "mutate: <amount> RSX market cap"},
				})
			}

			ctx = execute(ctx, []step.Interface{
				st012.Step{Value: 1e05 /***/, Comment: "mutate: buy RSX for 100k DAI from pool"},
				st015.Step{Value: 1e05 /***/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
				st020.Step{ /***************/ Comment: "mutate: <amount> excess reserves in treasury"},
				st006.Step{ /***************/ Comment: "mutate: <amount> RSX circulating supply"},
				st007.Step{ /***************/ Comment: "mutate: <amount> RSX total supply"},
				st011.Step{ /***************/ Comment: "mutate: <amount> RSX market cap"},
			})
		}

		{
			backing.AddX(i)
			backing.AddY(ctx.Treasury.DAI.Backing, ctx.Pool.RSX.MarketCap)

			excess.AddX(i)
			excess.AddY(ctx.Treasury.DAI.Excess)

			price.AddX(i)
			price.AddY(ctx.RSX.Price.Floor, ctx.RSX.Price.Ceiling)

			supply.AddX(i)
			supply.AddY(ctx.Treasury.RSX.Supply.Circulating, ctx.Treasury.RSX.Supply.Total)
		}
	}

	{
		chs = append(chs, backing.Line())
		chs = append(chs, excess.Line())
		chs = append(chs, price.Line())
		chs = append(chs, supply.Line())
	}

	return chs, nil
}

func execute(ctx context.Context, stp []step.Interface) context.Context {
	var err error

	for _, s := range stp {
		ctx, err = s.Run(ctx)
		if err != nil {
			panic(err)
		}
	}

	return ctx
}

func backingCanIncrease(ctx context.Context, floor float64, growth float64) bool {
	// Calculate the amount of capital we can add to the price floor.
	add := ctx.Treasury.DAI.Excess / ctx.Treasury.RSX.Supply.Circulating

	// Calculate the amount of price floor increase we plan to do.
	gro := floor * growth

	// As long as the available capital to increase the price floor is smaller
	// than the increase we want to achieve, we cannot increase the backing.
	return add >= gro
}
