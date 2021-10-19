package chart

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Chart struct {
	title string
	xAxis []string
	yAxis YAxis
}

type YAxis struct {
	name   string
	values []opts.LineData
}

func New(t string) *Chart {
	c := &Chart{
		title: t,
	}

	return c
}

func (c *Chart) AddX(x int) {
	c.xAxis = append(c.xAxis, fmt.Sprint(x))
}

func (c *Chart) AddY(y float64) {
	c.yAxis.values = append(c.yAxis.values, opts.LineData{Value: y})
}

func (c *Chart) Line() *charts.Line {
	l := charts.NewLine()

	l.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: c.title}),
		charts.WithTooltipOpts(opts.Tooltip{AxisPointer: &opts.AxisPointer{Type: "cross"}, Show: true, Trigger: "axis"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Time"}),
		charts.WithYAxisOpts(opts.YAxis{Name: "DAI"}),
	)

	l.SetXAxis(c.xAxis)
	l.AddSeries(c.title, c.yAxis.values)

	l.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	return l
}
