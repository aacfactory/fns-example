// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package examples

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/examples/components"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/logs"
	"github.com/aacfactory/fns/runtime"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/caches"
	"github.com/aacfactory/fns/services/documents"
	"github.com/aacfactory/fns/services/validators"
)

var (
	_endpointName = []byte("examples")
	_helloFnName  = []byte("hello")
)

func Hello(ctx context.Context, param HelloParam) (result HelloResults, err error) {
	// validate param
	if err = validators.Validate(param); err != nil {
		return
	}
	// cache
	cached, cacheExist, cacheGetErr := caches.Get(ctx, &param)
	if cacheGetErr != nil {
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(cacheGetErr).With("fns", "caches").Message("fns: get cache failed")
		}
	}
	if cacheExist {
		response := services.NewResponse(cached)
		result, err = services.ValueOfResponse[HelloResults](response)
		if err == nil {
			return
		}
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(cacheGetErr).With("fns", "caches").Message("fns: scan cached value failed")
		}
	}
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _helloFnName, param)
	if handleErr != nil {
		err = handleErr
		return
	}
	result, err = services.ValueOfResponse[HelloResults](response)
	//scanErr := response.Scan(&result)
	//if scanErr != nil {
	//	err = scanErr
	//	return
	//}
	return
}

func _helloFn(ctx services.Request) (v interface{}, err error) {
	// param
	param, paramErr := services.ValueOfParam[HelloParam](ctx.Param())
	if paramErr != nil {
		err = errors.BadRequest("scan params failed").WithCause(paramErr)
		return
	}
	// validate
	if err = validators.Validate(param); err != nil {
		return
	}
	// cache
	cached, cacheExist, cacheGetErr := caches.Get(ctx, param)
	if cacheGetErr != nil {
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(cacheGetErr).With("fns", "caches").Message("fns: get cache failed")
		}
	}
	if cacheExist {
		v = cached
		return
	}
	// handle
	v, err = hello(ctx, param)
	return
}

func Component[C services.Component](ctx context.Context, name string) (component C, has bool) {
	component, has = services.LoadComponent[C](ctx, _endpointName, name)
	return
}

func Service() (v services.Service) {
	v = &_service{
		Abstract: services.NewAbstract(
			string(_endpointName),
			false,
			&components.HelloComponent{},
		),
	}
	return
}

type _service struct {
	services.Abstract
}

func (svc *_service) Construct(options services.Options) (err error) {
	if err = svc.Abstract.Construct(options); err != nil {
		return
	}
	return
}

func (svc *_service) Document() (document documents.Endpoint) {
	document.AddFn(
		documents.NewFn(string(_helloFnName)).
			SetInfo("", "").
			SetReadonly(false).SetDeprecated(false).
			SetAuthorization(false).SetPermission(false).
			SetParam(documents.Unknown()).
			SetResult(documents.Unknown()).
			SetErrors("s"),
	)
	document.AddFn(
		documents.NewFn(string(_helloFnName)).
			SetInfo("", "").
			SetReadonly(false).SetDeprecated(false).
			SetAuthorization(false).SetPermission(false).
			SetParam(documents.Unknown()).
			SetResult(documents.Unknown()).
			SetErrors("s"),
	)
	return
}
