package simulation

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidArgumentError = &tracer.Error{
	Kind: "invalidArgumentError",
}

func IsInvalidArgument(err error) bool {
	return errors.Is(err, invalidArgumentError)
}

var invalidConfigError = &tracer.Error{
	Kind: "invalidConfigError",
}

func IsInvalidConfig(err error) bool {
	return errors.Is(err, invalidConfigError)
}
