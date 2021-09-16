package posts

import (
	"fmt"
	"github.com/aacfactory/configuares"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

const (
	Namespace = "posts"

	CreateFn = "create"
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
