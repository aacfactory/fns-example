package users

import (
	"fmt"
	"github.com/aacfactory/configuares"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

const (
	Namespace = "users"

	CreateFn = "create"
	GetFn    = "get"
	ListFn   = "list"

	LoginFn = "login"
	RevokeTokenFn = "revoke_token"
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
		err = svc.invokeCreateFn(context, argument)
	case ListFn:
		result, err = svc.invokeListFn(context, argument)
	case GetFn:
		result, err = svc.invokeGetFn(context, argument)
	case LoginFn:
		result, err = svc.invokeLoginFn(context, argument)
	case RevokeTokenFn:
		err = svc.invokeTokenRevokeFn(context)
	default:
		err = errors.NotFound(fmt.Sprintf("%s was not found in %s", fn, Namespace))
	}
	return
}

func (svc *_service) Close() (err error) {

	return
}

func (svc *_service) invokeCreateFn(context fns.Context, argument fns.Argument) (err errors.CodeError) {
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
	err = create(context, param)
	return
}

func (svc *_service) invokeGetFn(context fns.Context, argument fns.Argument) (result interface{}, err errors.CodeError) {
	context = fns.WithFn(context, GetFn)

	param := GetParam{}
	argErr := argument.As(&param)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("fns %s/%s: parse argument failed", Namespace, GetFn)).WithCause(argErr)
		return
	}

	// validate arg
	validateErr := context.App().Validate(param)
	if validateErr != nil {
		err = validateErr
		return
	}
	// handle fn
	result, err = get(context, param)
	return
}

func (svc *_service) invokeListFn(context fns.Context, argument fns.Argument) (result interface{}, err errors.CodeError) {
	context = fns.WithFn(context, ListFn)

	param := ListParam{}
	argErr := argument.As(&param)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("fns %s/%s: parse argument failed", Namespace, ListFn)).WithCause(argErr)
		return
	}

	// validate arg
	validateErr := context.App().Validate(param)
	if validateErr != nil {
		err = validateErr
		return
	}
	// handle fn
	result, err = list(context, param)
	return
}

func (svc *_service) invokeLoginFn(context fns.Context, argument fns.Argument) (result interface{}, err errors.CodeError) {
	context = fns.WithFn(context, LoginFn)

	param := LoginParam{}
	argErr := argument.As(&param)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("fns %s/%s: parse argument failed", Namespace, LoginFn)).WithCause(argErr)
		return
	}

	// validate arg
	validateErr := context.App().Validate(param)
	if validateErr != nil {
		err = validateErr
		return
	}
	// handle fn
	result, err = login(context, param)
	return
}

func (svc *_service) invokeTokenRevokeFn(context fns.Context) ( err errors.CodeError) {
	context = fns.WithFn(context, RevokeTokenFn)

	authorization, hasAuthorization := context.User().Authorization()
	if !hasAuthorization {
		err = errors.Unauthorized(fmt.Sprintf("fns %s/%s: no authorization", Namespace, RevokeTokenFn))
		return
	}
	authorizationErr := context.App().Authorizations().Decode(context, authorization)
	if authorizationErr != nil {
		err = errors.Unauthorized(fmt.Sprintf("fns %s/%s: invalid authorization", Namespace, RevokeTokenFn)).WithCause(authorizationErr)
		return
	}
	// handle fn
	err = tokenRevoke(context)
	return
}
