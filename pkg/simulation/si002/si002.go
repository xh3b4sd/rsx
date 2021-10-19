package si002

import (
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
	var chs []*charts.Line

	var excess *chart.Chart
	{
		excess = chart.New("Excess Reserves")
	}

	var backing *chart.Chart
	{
		backing = chart.New("DAI Backing")
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
		ctx = execute(ctx, []step.Interface{
			st012.Step{Value: 1e05 /***/, Comment: "mutate: buy RSX for 100k DAI from pool"},
			st015.Step{Value: 1e05 /***/, Comment: "mutate: arb RSX for 100k DAI between protocol and pool"},
			st020.Step{ /***************/ Comment: "mutate: <amount> excess reserves in treasury"},
		})

		{
			backing.AddX(i)
			backing.AddY(ctx.Treasury.DAI.Backing)

			excess.AddX(i)
			excess.AddY(ctx.Treasury.DAI.Excess)
		}
	}

	{
		chs = append(chs, excess.Line())
		chs = append(chs, backing.Line())
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
