package hellos

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/hellos/components"
	"github.com/aacfactory/fns/service"
)

type HelloArgument struct {
	Name string `json:"name" validate:"required" message:"name is invalid"`
}

type HelloResult struct {
	World string `json:"world"`
}

func hello(ctx context.Context, argument HelloArgument) (result *HelloResult, err errors.CodeError) {
	logger := service.GetLog(ctx)
	if logger.DebugEnabled() {
		logger.Debug().Caller().Message("handling...")
	}
	name := argument.Name
	if name == "error" {
		err = errors.ServiceError("service: service error")
		return
	}
	world := components.GetWorld(ctx)
	result = &HelloResult{
		World: fmt.Sprintf("%s: %s", world.Value(), argument.Name),
	}
	return
}
