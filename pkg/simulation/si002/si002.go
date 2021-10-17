package si002

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/xh3b4sd/tracer"
)

const (
	htmlFile = ".render/html/si002.html"
)

func Run() error {
	{
		p := filepath.Dir(htmlFile)

		err := os.MkdirAll(p, os.ModePerm)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		p := components.NewPage()

		p.AddCharts(
			lineSmooth(),
		)

		f, err := os.Create(htmlFile)
		if err != nil {
			return tracer.Mask(err)
		}

		p.Render(io.MultiWriter(f))
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

func lineSmooth() *charts.Line {
	l := charts.NewLine()

	l.SetGlobalOptions(charts.WithTitleOpts(opts.Title{Title: "foo"}))

	l.SetXAxis([]string{"1", "2", "3"})
	l.AddSeries("A", []opts.LineData{{Value: 5}, {Value: 3}, {Value: 4}})
	l.AddSeries("B", []opts.LineData{{Value: 4}, {Value: 6}, {Value: 2}})

	l.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	return l
}
