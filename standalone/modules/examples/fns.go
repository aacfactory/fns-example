// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package examples

import (
	"context"
	"time"

	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/examples/components"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/documents"
)

const (
	_name    = "examples"
	_helloFn = "hello"
)

func Hello(ctx context.Context, argument HelloArgument) (result HelloArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("examples: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _helloFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("examples: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _helloFn).
			WithCause(scanErr)
		return
	}
	return
}

func Service() (v service.Service) {
	v = &_service_{
		Abstract: service.NewAbstract(
			_name,
			false,
			[]service.Component{
				&components.HelloComponent{},
			}...,
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
		// param
		param := HelloArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("examples: decode request argument failed").WithCause(paramErr)
			break
		}
		// make timeout context
		var cancel context.CancelFunc = nil
		ctx, cancel = context.WithTimeout(ctx, time.Duration(1000000000))
		// barrier
		v, err = svc.Barrier(ctx, _helloFn, argument, func() (v interface{}, err errors.CodeError) {
			// execute function
			v, err = hello(ctx, param)
			return
		})
		// cancel timeout context
		cancel()
		break
	default:
		err = errors.Warning("examples: fn was not found").WithMeta("service", _name).WithMeta("fn", fn)
		break
	}
	return
}

func (svc *_service_) Document() (doc service.Document) {
	document := documents.NewService(_name, "Example service")
	// hello
	document.AddFn(
		"hello", "Hello", "Hello", false, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/examples", "HelloArgument").
			SetTitle("Hello function argument").
			SetDescription("Hello function argument").
			AddProperty(
				"world",
				documents.String().
					SetTitle("Name").
					SetDescription("Name").
					AsRequired().
					SetValidation(documents.NewElementValidation("world_required", "zh", "世界是必须的", "en", "world is required")),
			),
		documents.Array(documents.String()).
			SetPath("github.com/aacfactory/fns-example/standalone/modules/examples").
			SetName("HelloResults").
			SetTitle("Hello Results").
			SetDescription("Hello Results"),
		[]documents.FnError{
			{
				Name_: "examples_hello_failed",
				Descriptions_: map[string]string{
					"zh": "错误",
					"en": "failed",
				},
			},
		},
	)

	doc = document
	return

}
