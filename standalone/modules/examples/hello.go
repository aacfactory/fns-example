package examples

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
)

// HelloArgument
// @title Hello function argument
// @description Hello function argument
type HelloArgument struct {
	// World
	// @title Name
	// @description Name
	// @validate-message-i18n >>>
	// zh: 世界是必须的
	// en: world is required
	// <<<
	World string `json:"world" validate:"required" validate-message:"world_required"`
}

// HelloResults
// @title Hello Results
// @description Hello Results
type HelloResults []string

// hello
// @fn hello
// @timeout 1s
// @barrier
// @title Hello
// @errors >>>
// + examples_hello_failed
// 	- zh: 错误
//	- en: failed
// <<<
// @description >>>
// Hello
// <<<
func hello(ctx context.Context, argument HelloArgument) (result HelloResults, err errors.CodeError) {
	if argument.World == "error" {
		err = errors.ServiceError("examples_hello_failed")
		return
	}
	result = HelloResults{fmt.Sprintf("hello %s!", argument.World)}
	return
}
