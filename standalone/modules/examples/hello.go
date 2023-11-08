package examples

import (
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns/context"
)

// HelloParam
// @title Hello function argument
// @description Hello function argument
type HelloParam struct {
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
//   - zh: 错误
//   - en: failed
//
// <<<
// @description >>>
// Hello
// <<<
func hello(ctx context.Context, param HelloParam) (result HelloResults, err error) {
	if param.World == "error" {
		err = errors.ServiceError("examples_hello_failed")
		return
	}

	result = HelloResults{fmt.Sprintf("hello %s!", param.World)}
	return
}
