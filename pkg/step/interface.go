package step

import "github.com/xh3b4sd/rsx/pkg/context"

type Interface interface {
	Com() string
	Run(ctx context.Context) (context.Context, error)
}
