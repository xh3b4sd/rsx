package chart

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Chart struct {
	title string
	xAxis XAxis
	yAxis []YAxis
}

type XAxis struct {
	values []string
}

type YAxis struct {
	max    int
	min    int
	name   string
	values []opts.LineData
}

func New(t string, n ...string) *Chart {
	c := &Chart{
		title: t,
	}

	for _, v := range n {
		c.yAxis = append(c.yAxis, YAxis{name: v})
	}

	if len(c.yAxis) == 0 {
		c.yAxis = append(c.yAxis, YAxis{name: t})
	}

	return c
}

func (c *Chart) AddX(x int) {
	c.xAxis.values = append(c.xAxis.values, fmt.Sprint(x))
}

func (c *Chart) AddY(y ...float64) {
	for i, v := range y {
		c.yAxis[i].values = append(c.yAxis[i].values, opts.LineData{Value: v})
	}
}

func (c *Chart) Line() *charts.Line {
	l := charts.NewLine()

	l.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: c.title}),
		charts.WithTooltipOpts(opts.Tooltip{AxisPointer: &opts.AxisPointer{Type: "cross"}, Show: true, Trigger: "axis"}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "85", Top: "3"}),
		charts.WithXAxisOpts(opts.XAxis{Name: "Time"}),
		charts.WithYAxisOpts(opts.YAxis{Max: c.yAxis[0].max, Min: c.yAxis[0].min, Name: "DAI"}),
	)

	l.SetXAxis(c.xAxis.values)

	for i := range c.yAxis {
		l.AddSeries(c.yAxis[i].name, c.yAxis[i].values, charts.WithLineChartOpts(opts.LineChart{ShowSymbol: false}))
	}

	return l
}

func (c *Chart) MaxY(m int) {
	for i := range c.yAxis {
		c.yAxis[i].max = m
	}
}

func (c *Chart) MinY(m int) {
	for i := range c.yAxis {
		c.yAxis[i].min = m
	}
}
