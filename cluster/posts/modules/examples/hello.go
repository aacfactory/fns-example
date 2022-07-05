package examples

import (
	"context"
	"fmt"

	"github.com/aacfactory/errors"
)

// HelloArgument
// @title Hello Argument
// @description Hello Argument
type HelloArgument struct {
	// Name
	// @title Name
	// @description Name
	Name string `json:"name" message:"name is required" validate:"required"`
}

// HelloResult
// @title Hello Result
// @description Hello Result
type HelloResult struct {
	// Name
	// @title Name
	// @description Name
	Name string `json:"name"`
}

// hello
// @fn hello
// @validate true
// @authorization false
// @permission false
// @internal false
// @title Hello
// @description >>>
// Hello Fn
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | examples_hello_failed    | 500     | hello failed                  |
// <<<
func hello(ctx context.Context, argument HelloArgument) (result *HelloResult, err errors.CodeError) {
	if argument.Name == "error" {
		err = errors.ServiceError("examples_hello_failed")
		return
	}
	result = &HelloResult{
		Name: fmt.Sprintf("hello %s!", argument.Name),
	}
	return
}
