package st009

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var executionFailedError = &tracer.Error{
	Kind: "executionFailedError",
}

func IsExecutionFailed(err error) bool {
	return errors.Is(err, executionFailedError)
}
