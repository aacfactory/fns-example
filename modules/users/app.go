package users

import (
	"fmt"
	"github.com/aacfactory/configuares"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

const (
	namespace = "users"
)

type Service struct {
}

func (svc *Service) Namespace() string {
	return namespace
}

func (svc *Service) Build(config configuares.Config) (err error) {

	return
}

func (svc *Service) Description() (description []byte) {
	return
}

func (svc *Service) Handle(context fns.Context, fn string, argument fns.Argument) (result interface{}, err errors.CodeError) {
	switch fn {
	case "create":
		result, err = svc.invokeCreateFn(context, fn, argument)
	default:
		err = errors.NotFound(fmt.Sprintf("%s was not found in %s", fn, namespace))
	}
	return
}

func (svc *Service) Close() (err error) {

	return
}

func (svc *Service) invokeCreateFn(context fns.Context, fn string, argument fns.Argument) (result interface{}, err errors.CodeError) {
	// context with fn
	context = fns.WithFn(context, fn)
	// scan arg
	arg := UserCreateArg{}
	argErr := argument.As(&arg)
	if argErr != nil {
		err = errors.BadRequest(fmt.Sprintf("%s/%s: parse argument failed, %v", namespace, fn, argErr))
		return
	}
	// validate arg
	validateErr := context.Validate(arg)
	if validateErr != nil {
		err = validateErr
		return
	}
	// handle fn
	handleErr := create(context, arg)
	if handleErr != nil {
		err = handleErr
		return
	}
	result = fns.Empty{}
	return
}
