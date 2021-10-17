package simulation

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"

	"github.com/xh3b4sd/rsx/pkg/simulation/si001"
	"github.com/xh3b4sd/rsx/pkg/simulation/si002"
)

type runner struct {
	logger logger.Interface
}

func (r *runner) Run(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	err := r.run(ctx, cmd, args)
	if err != nil {
		return tracer.Mask(err)
	}

	return nil
}

func (r *runner) run(ctx context.Context, cmd *cobra.Command, args []string) error {
	// We expect exactly one argument, like 001 to run the first simulation.
	if len(args) != 1 {
		return tracer.Maskf(invalidArgumentError, "simulation not provided as argument")
	}

	switch args[0] {
	case "001":
		err := si001.Run()
		if err != nil {
			return tracer.Mask(err)
		}
	case "002":
		err := si002.Run()
		if err != nil {
			return tracer.Mask(err)
		}
	default:
		return tracer.Maskf(invalidArgumentError, "simulation %s not registered", args[0])
	}

	return nil
}
