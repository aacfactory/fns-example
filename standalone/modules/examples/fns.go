// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package examples

import (
	"github.com/aacfactory/fns/commons/bytex"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/runtime"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/validators"

	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/examples/components"
)

const (
	_name    = "examples"
	_helloFn = "hello"
)

func Hello(ctx context.Context, param HelloParam) (result HelloResults, err error) {

	return
}

func Service() (v services.Service) {
	v = &service{
		Abstract: services.NewAbstract(
			_name,
			false,
			[]services.Component{
				&components.HelloComponent{},
			}...,
		),
	}
	return
}

type service struct {
	services.Abstract
}

func (svc *service) Handle(ctx services.Request) (v interface{}, err error) {
	_, fn := ctx.Fn()
	fnName := bytex.ToString(fn)
	switch fnName {
	case _helloFn:
		// barrier
		key := ctx.Hash()
		if len(ctx.Header().Token()) > 0 {
			key = append(key, ctx.Header().Token()...)
		}
		v, err = runtime.Barrier(ctx, key, func() (v interface{}, err error) {
			// authorizations
			// todo

			// permissions
			// todo

			// param
			param := HelloParam{}
			if paramErr := ctx.Argument().As(&param); paramErr != nil {
				err = errors.BadRequest("invalid body").WithMeta("service", svc.Name()).WithMeta("fn", fnName)
				return
			}
			// validate param
			if validateErr := validators.Validate(param); validateErr != nil {
				err = errors.BadRequest("invalid body").WithCause(validateErr).WithMeta("service", svc.Name()).WithMeta("fn", fnName)
				return
			}
			// handle
			v, err = hello(ctx, param)
			if err != nil {
				err = errors.Map(err).WithMeta("service", svc.Name()).WithMeta("fn", fnName)
				return
			}
			return
		})
	default:
		err = errors.NotFound("fn was not found").WithMeta("service", svc.Name()).WithMeta("fn", fnName)
	}
	return
}
