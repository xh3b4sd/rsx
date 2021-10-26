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
	unit   string
	values []string
}

type YAxis struct {
	ceiling         float64
	floor           float64
	hasSeparateAxis bool
	max             float64
	min             int
	name            string
	rfv             int
	values          []opts.LineData
	unit            string
}

func New(t string, n ...string) *Chart {
	c := &Chart{
		title: t,
		xAxis: XAxis{
			unit: "Time",
		},
	}

	for _, v := range n {
		c.yAxis = append(c.yAxis, YAxis{name: v})
	}

	if len(c.yAxis) == 0 {
		c.yAxis = append(c.yAxis, YAxis{name: t})
	}

	return c
}

func (c *Chart) AddX(x float64) {
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
		charts.WithXAxisOpts(opts.XAxis{Name: c.xAxis.unit}),
		charts.WithYAxisOpts(opts.YAxis{Max: c.yAxis[0].max, Min: c.yAxis[0].min, Name: c.yAxis[0].unit}),
	)

	l.SetXAxis(c.xAxis.values)

	for _, y := range c.yAxis {
		index := 0

		if y.hasSeparateAxis {
			index = 1
			l.ExtendYAxis(opts.YAxis{
				Name: y.name,
				Type: "value",
				Show: true,
				Min:  y.min,
				Max:  y.max,
			})
		}

		ops := []charts.SeriesOpts{
			charts.WithLineChartOpts(opts.LineChart{ShowSymbol: false, YAxisIndex: index}),
		}

		if y.ceiling != 0 && y.floor != 0 {
			ops = append(ops, charts.WithMarkLineNameYAxisItemOpts(
				opts.MarkLineNameYAxisItem{Name: "Ceiling", YAxis: y.ceiling},
			))
			ops = append(ops, charts.WithMarkLineNameYAxisItemOpts(
				opts.MarkLineNameYAxisItem{Name: "Floor", YAxis: y.floor},
			))
			ops = append(ops, charts.WithMarkLineStyleOpts(
				opts.MarkLineStyle{Label: &opts.Label{Show: true, Position: "end", Formatter: "{b}"}, Symbol: []string{"none", "none"}},
			))
		}

		if y.rfv != 0 {
			ops = append(ops, charts.WithMarkLineNameXAxisItemOpts(
				opts.MarkLineNameXAxisItem{Name: "RFV", XAxis: y.rfv},
			))
			ops = append(ops, charts.WithMarkLineStyleOpts(
				opts.MarkLineStyle{Label: &opts.Label{Show: false}, Symbol: []string{"none", "none"}},
			))
		}

		l.AddSeries(y.name, y.values, ops...)
	}

	return l
}

func (c *Chart) MaxY(m float64) {
	for i := range c.yAxis {
		c.yAxis[i].max = m
	}
}

func (c *Chart) MinY(m int) {
	for i := range c.yAxis {
		c.yAxis[i].min = m
	}
}

func (c *Chart) RFVY(i int, r int) {
	c.yAxis[i].rfv = r
}

func (c *Chart) SetCeilingAndFloor(cei float64, flo float64) {
	c.yAxis[0].ceiling = cei
	c.yAxis[0].floor = flo
}

func (c *Chart) SetSeparateYAxis(i int, n string, m float64) {
	c.yAxis[i].hasSeparateAxis = true
	c.yAxis[i].name = n
	c.yAxis[i].max = m
}

func (c *Chart) UnitX(u string) {
	c.xAxis.unit = u
}

func (c *Chart) UnitY(u string) {
	for i := range c.yAxis {
		c.yAxis[i].unit = u
	}
}
