package si001

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/bond"
	"github.com/xh3b4sd/rsx/pkg/chart"
	"github.com/xh3b4sd/rsx/pkg/context"
	"github.com/xh3b4sd/rsx/pkg/round"
	"github.com/xh3b4sd/rsx/pkg/step"
	"github.com/xh3b4sd/rsx/pkg/step/st002"
	"github.com/xh3b4sd/rsx/pkg/step/st003"
	"github.com/xh3b4sd/rsx/pkg/step/st006"
	"github.com/xh3b4sd/rsx/pkg/step/st007"
	"github.com/xh3b4sd/rsx/pkg/step/st011"
	"github.com/xh3b4sd/rsx/pkg/step/st015"
	"github.com/xh3b4sd/rsx/pkg/step/st020"
	"github.com/xh3b4sd/rsx/pkg/step/st021"
)

const (
	htmlFile = ".render/html/si001.html"
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
	// dao is the percentage of volume inflow diverted into the DAO holdings.
	var dao float64
	{
		dao = 0.10
	}

	// deb is the percentage of volume inflow diverted into paying back protocol
	// debt.
	// var deb float64
	// {
	// 	deb = 0.10
	// }

	var floor float64
	{
		floor = 0.10
	}

	// multiple is the multiple to be applied to floor in order to compute the
	// actual price floor. E.g. 2.00 means price ceiling is always 2x price
	// floor.
	var multiple float64
	{
		multiple = 5.00
	}

	var ceiling float64
	{
		ceiling = floor * multiple
	}

	// growth is the multiple at which price floor and price ceiling grow. E.g.
	// 0.10 means the price floor rises by 10% on each increase.
	var growth float64
	{
		growth = 0.20
	}

	// Track the ROI based on an initial position in the price ceiling.
	// Eventually the price floor grows beyond the initial position and creates
	// a positive return on investment.
	var pos float64
	var roi float64
	{
		pos = ceiling
	}

	var chs []*charts.Line

	var backing *chart.Chart
	{
		backing = chart.New("DAI Backing / Market Cap", "DAI Backing", "Market Cap")
		backing.MaxY(5e08)
		backing.MinY(0)
		backing.UnitY("DAI")
	}

	var bonding *chart.Chart
	{
		bonding = chart.New("Bond Discounts / Bond Volume", "RSX Price", "Bond Price")
		bonding.MaxY(6)
		bonding.MinY(0)
		bonding.UnitX("Volume")
		bonding.UnitY("Price")
	}

	var excess *chart.Chart
	{
		excess = chart.New("Excess Reserves / DAO Holdings", "Excess Reserves", "DAO Holdings")
		excess.MaxY(8e06)
		excess.MinY(0)
		excess.UnitY("DAI")
	}

	var price *chart.Chart
	{
		price = chart.New("Price Floor / Price Ceiling", "Price Floor", "Price Ceiling", "RFV Multiple")
		price.MaxY(15)
		price.MinY(0)
		price.UnitY("DAI")
	}

	ctx = execute(ctx, []step.Interface{
		st002.Step{Value: floor /*****/, Comment: fmt.Sprintf("mutate: set %.2f DAI price floor", floor)},
		st003.Step{Value: ceiling /***/, Comment: fmt.Sprintf("mutate: set %.2f DAI price ceiling", ceiling)},

		st006.Step{ /******************/ Comment: "mutate: <amount> RSX circulating supply"},
		st007.Step{ /******************/ Comment: "mutate: <amount> RSX total supply"},
		st011.Step{ /******************/ Comment: "mutate: <amount> RSX market cap"},
		st020.Step{ /******************/ Comment: "mutate: <amount> excess reserves in treasury"},
	})

	o := sync.Once{}

	var sho bool

	for i := 0; i < 1000; i++ {
		{
			if i != 0 && floorCanIncrease(ctx, floor, growth) {
				floor = floor * (1 + growth)
				ceiling = floor * multiple

				growth = growth * 0.95
				if growth < 0.01 {
					growth = 0.01
				}

				ste := []step.Interface{
					st002.Step{Value: floor /*****/, Comment: fmt.Sprintf("mutate: set %.2f DAI price floor", floor)},
					st003.Step{Value: ceiling /***/, Comment: fmt.Sprintf("mutate: set %.2f DAI price ceiling", ceiling)},
					st006.Step{ /******************/ Comment: /*********/ "mutate: <amount> RSX circulating supply"},
					st007.Step{ /******************/ Comment: /*********/ "mutate: <amount> RSX total supply"},
					st011.Step{ /******************/ Comment: /*********/ "mutate: <amount> RSX market cap"},
					st020.Step{ /******************/ Comment: /*********/ "mutate: <amount> excess reserves in treasury"},
				}

				ctx = execute(ctx, ste)
			}
		}

		{
			ste := []step.Interface{
				st015.Step{Value: 1e05 /***/, Comment: /*********/ "mutate: add 100k DAI treasury inflow"},
				st006.Step{ /***************/ Comment: /*********/ "mutate: <amount> RSX circulating supply"},
				st007.Step{ /***************/ Comment: /*********/ "mutate: <amount> RSX total supply"},
				st011.Step{ /***************/ Comment: /*********/ "mutate: <amount> RSX market cap"},
				st020.Step{ /***************/ Comment: /*********/ "mutate: <amount> excess reserves in treasury"},
				st021.Step{Value: dao /****/, Comment: fmt.Sprintf("mutate: add %.2f DAI to DAO", dao)},
			}

			ctx = execute(ctx, ste)
		}

		{
			o.Do(func() {
				ste := []step.Interface{
					st015.Step{Value: 1e07, Comment: "mutate: add 10M DAI treasury inflow"},
					st006.Step{ /*********/ Comment: "mutate: <value> RSX circulating supply"},
					st007.Step{ /*********/ Comment: "mutate: <value> RSX total supply"},
					st011.Step{ /*********/ Comment: "mutate: <value> RSX market cap"},
					st020.Step{ /*********/ Comment: "mutate: <value> excess reserves in treasury"},
				}

				ctx = execute(ctx, ste)
			})
		}

		{
			roi = floor / pos
			if !sho && roi >= 1 {
				sho = true
				price.RFVY(2, i)
			}
		}

		{
			backing.AddX(float64(i))
			excess.AddX(float64(i))
			price.AddX(float64(i))

			backing.AddY(ctx.Treasury.DAI.Backing, ctx.Treasury.RSX.Supply.MarketCap)
			excess.AddY(ctx.Treasury.DAI.Excess, ctx.Treasury.DAI.DAO/2) // divided by 2 for innovation fund
			price.AddY(ctx.RSX.Price.Floor, ctx.RSX.Price.Ceiling, roi)
		}
	}

	{
		c := float64(5)
		f := float64(1)
		s := 1000

		bonding.SetCeilingAndFloor(c, f)

		for i := 0; i < s; i++ {
			p := ((c + 1) / float64(s) * float64(i))

			d := bond.Discount(p, f, c)
			v := bond.Volume(p, f, c)

			bonding.AddX(round.RoundP(v, 0))
			bonding.AddY(round.RoundP(p, 3), round.RoundP(d, 3))
		}
	}

	{
		chs = append(chs, backing.Line())
		chs = append(chs, excess.Line())
		chs = append(chs, price.Line())
		chs = append(chs, bonding.Line())
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

func floorCanIncrease(ctx context.Context, floor float64, growth float64) bool {
	// Calculate the amount of capital we can add to the price floor.
	add := ctx.Treasury.DAI.Excess / ctx.Treasury.RSX.Supply.Circulating

	// Calculate the amount of price floor increase we plan to do.
	gro := floor * growth

	// As long as the available capital to increase the price floor is smaller
	// than the increase we want to achieve, we cannot increase the backing.
	return add >= gro
}
