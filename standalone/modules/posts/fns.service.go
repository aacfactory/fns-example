package posts

import (
	"fmt"
	"github.com/aacfactory/configuares"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/redis"
	"github.com/aacfactory/json"
	"time"
)

const (
	Namespace = "posts"

	CreateFn = "create"
	GetFn    = "get"
)

func Service() fns.Service {
	return &_service{}
}

type _service struct {
}

func (svc *_service) Namespace() string {
	return Namespace
}

func (svc *_service) Internal() bool {
	return false
}

func (svc *_service) Build(_ configuares.Config) (err error) {

	return
}

func (svc *_service) Description() (description []byte) {
	return
}

func (svc *_service) Handle(context fns.Context, fn string, argument fns.Argument) (result interface{}, err errors.CodeError) {
	switch fn {
	case CreateFn:
		result, err = svc.invokeCreateFn(context, argument)
	case GetFn:
		result, err = svc.invokeGetFn(context, argument)
	default:
		err = errors.NotFound(fmt.Sprintf("%s was not found in %s", fn, Namespace))
	}
	return
}

func (svc *_service) Close() (err error) {

	return
}

func (svc *_service) invokeCreateFn(context fns.Context, argument fns.Argument) (v interface{}, err errors.CodeError) {
	// context with fn
	context = fns.WithFn(context, CreateFn)
	// check authorization
	// check use is active
	// check permissions

	// scan arg
	param := CreateParam{}
	argErr := argument.As(&param)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("fns %s/%s: parse argument failed", Namespace, CreateFn)).WithCause(argErr)
		return
	}
	// validate arg
	validateErr := context.App().Validate(param)
	if validateErr != nil {
		err = validateErr
		return
	}
	// handle fn
	v, err = create(context, param)
	return
}

func (svc *_service) invokeGetFn(context fns.Context, argument fns.Argument) (v interface{}, err errors.CodeError) {
	// context with fn
	context = fns.WithFn(context, CreateFn)
	// check authorization
	// check use is active
	// check permissions

	// scan arg
	param := GetParam{}
	argErr := argument.As(&param)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("fns %s/%s: parse argument failed", Namespace, CreateFn)).WithCause(argErr)
		return
	}
	// validate arg
	validateErr := context.App().Validate(param)
	if validateErr != nil {
		err = validateErr
		return
	}
	// cache
	key := param.Id
	result, cacheErr := redis.CacheGetWithSingleFight(context, key, 12*time.Second, func() (result json.RawMessage, err errors.CodeError) {
		// handle fn
		handleResult, handleErr := get(context, param)
		if handleErr != nil {
			err = handleErr
			return
		}
		p, encodeErr := json.Marshal(handleResult)
		if encodeErr != nil {
			err = errors.ServiceError("encode json failed").WithCause(encodeErr)
			return
		}
		result = p
		return
	})
	if cacheErr != nil {
		err = cacheErr
		return
	}
	v = result
	return
}
