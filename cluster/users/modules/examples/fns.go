// NOTE: This file has been automatically generated by github.com/aacfactory/gcg. Dont Edit it.

package examples

import (
	"context"

	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/documents"
	"github.com/aacfactory/fns/service/validators"
)

const (
	_name    = "examples"
	_helloFn = "hello"
)

func Hello(ctx context.Context, argument HelloArgument) (result *HelloResult, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.NotFound("endpoint was not found").WithMeta("name", _name)
		return
	}
	fr := endpoint.Request(ctx, _helloFn, service.NewArgument(argument))
	handled := HelloResult{}
	hasResult, handleErr := fr.Get(ctx, &handled)
	if handleErr != nil {
		err = handleErr
		return
	}
	if hasResult {
		result = &handled
	}
	return
}

func Service() (svc service.Service) {
	components := []service.Component{}
	svc = &_service_{
		Abstract: service.NewAbstract(
			_name,
			false,
			components...,
		),
	}
	return
}

type _service_ struct {
	service.Abstract
}

func (svc *_service_) Handle(ctx context.Context, fn string, argument service.Argument) (v interface{}, err errors.CodeError) {
	switch fn {
	case _helloFn:
		// argument
		arg := HelloArgument{}
		scanErr := argument.As(&arg)
		if scanErr != nil {
			err = errors.BadRequest("examples: scan request argument failed").WithCause(scanErr).WithMeta("service", _name).WithMeta("fn", _helloFn)
			return
		}
		validateErr := validators.Validate(arg)
		if validateErr != nil {
			err = errors.BadRequest("examples: invalid request argument").WithMeta("service", _name).WithMeta("fn", _helloFn).WithCause(validateErr)
			return
		}
		// handle
		v, err = hello(ctx, arg)
		break
	default:
		err = errors.NotFound("examples: fn was not found").WithMeta("service", _name).WithMeta("fn", fn)
		break
	}
	return

}

func (svc *_service_) Document() (doc service.Document) {
	sd := documents.NewService(_name, "Example service")
	sd.AddFn(
		"hello", "Hello", "Hello Fn\n----------\nerrors:\n| Name                     | Code    | Description                   |\n|--------------------------|---------|-------------------------------|\n| examples_hello_failed    | 500     | hello failed                  |", false, false,
		documents.Struct("main/modules/examples", "HelloArgument", "Hello Argument", "Hello Argument").
			AddProperty("name",
				documents.String().SetTitle("Name").SetDescription("Name").AsRequired(`validate:"required" message:"name is required"`),
			),
		documents.Struct("main/modules/examples", "HelloResult", "Hello Result", "Hello Result").
			AddProperty("name",
				documents.String().SetTitle("Name").SetDescription("Name"),
			),
	)
	doc = sd
	return
}
