package st002

import "github.com/xh3b4sd/rsx/pkg/context"

type Step struct {
	Comment string
	Value   float64
}

func (s Step) Com() string {
	return s.Comment
}

// set price floor to <Value>
func (s Step) Run(ctx context.Context) (context.Context, error) {
	ctx.RSX.Price.Floor = s.Value

	return ctx, nil
}
