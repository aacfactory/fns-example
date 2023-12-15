// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/users/middles"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/logs"
	"github.com/aacfactory/fns/runtime"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/caches"
	"github.com/aacfactory/fns/services/commons"
	"github.com/aacfactory/fns/services/documents"
	"github.com/aacfactory/fns/services/validators"
	"github.com/aacfactory/fns/transports/middlewares/cachecontrol"
)

var (
	_endpointName = []byte("users")
	_addFnName    = []byte("add")
	_getFnName    = []byte("get")
	_listFnName   = []byte("list")
)

// +-------------------------------------------------------------------------------------------------------------------+

func Add(ctx context.Context, param AddParam) (result User, err error) {
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _addFnName, param)
	if handleErr != nil {
		err = handleErr
		return
	}
	result, err = services.ValueOfResponse[User](response)
	return

}

func _add(ctx services.Request) (v any, err error) {
	// param
	param, paramErr := services.ValueOfParam[AddParam](ctx.Param())
	if paramErr != nil {
		err = errors.BadRequest("scan params failed").WithCause(paramErr)
		return
	}
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// handle
	v, err = add(ctx, param)
	return

}

// +-------------------------------------------------------------------------------------------------------------------+

func Get(ctx context.Context, param GetParam) (result User, err error) {
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// cache get
	cached, cacheExist, cacheGetErr := caches.Get(ctx, &param)
	if cacheGetErr != nil {
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(cacheGetErr).With("fns", "caches").Message("fns: get cache failed")
		}
	}
	if cacheExist {
		response := services.NewResponse(cached)
		result, err = services.ValueOfResponse[User](response)
		if err == nil {
			return
		}
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(err).With("fns", "caches").Message("fns: scan cached value failed")
		}
	}
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _getFnName, param)
	if handleErr != nil {
		err = handleErr
		return
	}
	result, err = services.ValueOfResponse[User](response)
	return

}

func _get(ctx services.Request) (v any, err error) {
	// param
	param, paramErr := services.ValueOfParam[GetParam](ctx.Param())
	if paramErr != nil {
		err = errors.BadRequest("scan params failed").WithCause(paramErr)
		return
	}
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// cache get
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
	v, err = get(ctx, param)
	// cache set
	if cacheSetErr := caches.Set(ctx, param, v, 10); cacheSetErr != nil {
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(cacheSetErr).With("fns", "caches").Message("fns: set cache failed")
		}
	}
	// cache control
	cachecontrol.Make(ctx, cachecontrol.MaxAge(10), cachecontrol.Public())
	return

}

// +-------------------------------------------------------------------------------------------------------------------+

func List(ctx context.Context) (result Users, err error) {
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _listFnName, nil)
	if handleErr != nil {
		err = handleErr
		return
	}
	result, err = services.ValueOfResponse[Users](response)
	return

}

func _list(ctx services.Request) (v any, err error) {
	// handle
	v, err = list(ctx)
	// cache control
	cachecontrol.Make(ctx, cachecontrol.MaxAge(10), cachecontrol.Public())
	return

}

// +-------------------------------------------------------------------------------------------------------------------+

func Component[C services.Component](ctx context.Context, name string) (component C, has bool) {
	component, has = services.LoadComponent[C](ctx, _endpointName, name)
	return
}

// +-------------------------------------------------------------------------------------------------------------------+

func Service() (v services.Service) {
	v = &_service{
		Abstract: services.NewAbstract(
			string(_endpointName),
			false,
		),
	}
	return
}

// +-------------------------------------------------------------------------------------------------------------------+

type _service struct {
	services.Abstract
}

func (svc *_service) Construct(options services.Options) (err error) {
	if err = svc.Abstract.Construct(options); err != nil {
		return
	}
	svc.AddFunction(commons.NewFn(string(_addFnName), false, false, true, false, false, false, _add))
	svc.AddFunction(commons.NewFn(string(_getFnName), true, false, true, true, true, true, _get))
	svc.AddFunction(commons.NewFn(string(_listFnName), true, false, true, true, true, true, _list, &Middle{}, &middles.Middle{}))
	return
}

func (svc *_service) Document() (document documents.Endpoint) {
	document = documents.New(svc.Name(), "Users", "Users")
	// add
	document.AddFn(
		documents.NewFn("add").
			SetInfo("add", "add user").
			SetReadonly(false).SetInternal(false).SetDeprecated(false).
			SetAuthorization(true).SetPermission(false).
			SetParam(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "AddParam").
				SetTitle("add param").
				SetDescription("add user param").
				AddProperty(
					"id",
					documents.String().
						SetTitle("user id").
						SetDescription("user id").
						AsRequired().
						SetValidation(documents.NewValidation("")),
				).
				AddProperty(
					"name",
					documents.String().
						SetTitle("name").
						SetDescription("name").
						AsRequired().
						SetValidation(documents.NewValidation("")),
				).
				AddProperty(
					"age",
					documents.Int64().
						SetTitle("age").
						SetDescription("age"),
				).
				AddProperty(
					"birthday",
					documents.Time().
						SetTitle("birthday").
						SetDescription("birthday"),
				)).
			SetResult(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "User").
				AddProperty(
					"id",
					documents.String(),
				).
				AddProperty(
					"name",
					documents.String(),
				).
				AddProperty(
					"age",
					documents.String(),
				).
				AddProperty(
					"birthday",
					documents.DateTime(),
				)).
			SetErrors("user_not_found\nzh: zh_message\nen: en_message"),
	)

	// get
	document.AddFn(
		documents.NewFn("get").
			SetInfo("get", "dafasdf\nadsfasfd").
			SetReadonly(true).SetInternal(false).SetDeprecated(false).
			SetAuthorization(true).SetPermission(true).
			SetParam(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "GetParam").
				SetTitle("get param").
				SetDescription("get param").
				AddProperty(
					"id",
					documents.String().
						SetTitle("id").
						SetDescription("id").
						SetValidation(documents.NewValidation("invalid_id").AddI18n("zh", "zh_message").AddI18n("en", "en_message")),
				)).
			SetResult(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "User").
				AddProperty(
					"id",
					documents.String(),
				).
				AddProperty(
					"name",
					documents.String(),
				).
				AddProperty(
					"age",
					documents.String(),
				).
				AddProperty(
					"birthday",
					documents.DateTime(),
				)).
			SetErrors("user_not_found\nzh: zh_message\nen: en_message"),
	)

	// list
	document.AddFn(
		documents.NewFn("list").
			SetInfo("list", "dafasdf\nadsfasfd").
			SetReadonly(true).SetInternal(false).SetDeprecated(false).
			SetAuthorization(true).SetPermission(true).
			SetParam(documents.Nil()).
			SetResult(documents.Array(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "User").
				AddProperty(
					"id",
					documents.String(),
				).
				AddProperty(
					"name",
					documents.String(),
				).
				AddProperty(
					"age",
					documents.String(),
				).
				AddProperty(
					"birthday",
					documents.DateTime(),
				)).
				SetPath("github.com/aacfactory/fns-example/standalone/modules/users").
				SetName("Users")).
			SetErrors("user_not_found\nzh: zh_message\nen: en_message"),
	)
	return
}
