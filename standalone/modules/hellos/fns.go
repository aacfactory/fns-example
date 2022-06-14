package hellos

import (
	"context"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/hellos/components"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/validators"
)

func Hello(ctx context.Context, argument HelloArgument) (result *HelloResult, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, "hellos")
	if !hasEndpoint {
		err = errors.NotFound("endpoint was not found").WithMeta("name", "hellos")
		return
	}
	fr := endpoint.Request(ctx, "hello", service.NewArgument(argument))
	handled := &HelloResult{}
	hasResult, handleErr := fr.Get(ctx, handled)
	if handleErr != nil {
		err = handleErr
		return
	}
	if hasResult {
		result = handled
	}
	return
}

func Service() (svc service.Service) {
	svc = &_service_{
		service.NewAbstract(
			"hellos",
			false,
			components.NewWorld(),
		),
	}
	return
}

type _service_ struct {
	service.Abstract
}

func (svc *_service_) Handle(ctx context.Context, fn string, argument service.Argument) (v interface{}, err errors.CodeError) {
	switch fn {
	case "hello":
		helloArgument := HelloArgument{}
		argumentAsErr := argument.As(&helloArgument)
		if argumentAsErr != nil {
			err = errors.BadRequest("service: bad argument").WithCause(argumentAsErr).WithMeta("service", "hellos").WithMeta("fn", fn)
			return
		}
		validateErr := validators.Validate(helloArgument)
		if validateErr != nil {
			err = errors.BadRequest("service: invalid argument").WithMeta("service", "hellos").WithMeta("fn", fn).WithCause(validateErr)
			return
		}
		v, err = hello(ctx, helloArgument)
		break
	default:
		err = errors.NotFound("service: fn was not found").WithMeta("service", "hellos").WithMeta("fn", fn)
		break
	}
	return
}

func (svc *_service_) Close() {
	if svc.Log().DebugEnabled() {
		svc.Log().Debug().Caller().Message("service: close")
	}
	return
}

func (svc *_service_) Document() (doc service.Document) {

	return
}
