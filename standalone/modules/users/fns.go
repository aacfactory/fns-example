// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package users

import (
	"github.com/aacfactory/errors"
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
	_getFnName    = []byte("get")
)

// +-------------------------------------------------------------------------------------------------------------------+

func Get(ctx context.Context, param GetParam) (result GetParam, err error) {
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
		scanErr := response.Scan(&result)
		if scanErr == nil {
			return
		}
		log := logs.Load(ctx)
		if log.WarnEnabled() {
			log.Warn().Cause(scanErr).With("fns", "caches").Message("fns: scan cached value failed")
		}
	}
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _getFnName, param)
	if handleErr != nil {
		err = handleErr
		return
	}
	scanErr := response.Scan(&result)
	if scanErr != nil {
		err = scanErr
		return
	}
	return

}

func _get(ctx services.Request) (v any, err error) {
	// param
	param := GetParam{}
	if paramErr := ctx.Param().Scan(&param); paramErr != nil {
		err = errors.BadRequest("scan params failed").WithCause(paramErr)
		return
	}
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
	svc.AddFunction(commons.NewFn(string(_getFnName), true, false, true, true, true, true, _get))
	return
}

func (svc *_service) Document() (document documents.Endpoint) {
	document = documents.New(svc.Name(), "Users", "Users")
	// get
	document.AddFn(
		documents.NewFn("get").
			SetInfo("get", "dafasdf\nadsfasfd").
			SetReadonly(true).SetDeprecated(false).
			SetAuthorization(true).SetPermission(true).
			SetParam(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "GetParam").
				SetTitle("get").
				SetDescription("get").
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
	return
}
