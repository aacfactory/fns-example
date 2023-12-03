package examples

import (
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/commons"
	"github.com/aacfactory/fns/transports"
	"sync/atomic"
	"time"
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

func (param *HelloParam) CacheKey(ctx context.Context) (key []byte, err error) {

	return
}

// HelloResults
// @title Hello Results
// @description Hello Results
type HelloResults []string

var (
	counter = atomic.Int64{}
)

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
	counter.Add(1)
	time.Sleep(500 * time.Millisecond)
	result = HelloResults{fmt.Sprintf("hello %s!", param.World)}
	fmt.Println(counter.Load())
	w := transports.LoadResponseWriter(ctx)
	fmt.Println(w.LocalValue([]byte("@fns:context:runtime")) != nil)
	return
}

type HelloMiddleware struct {
}

func (middle *HelloMiddleware) Handler(next commons.FnHandler) commons.FnHandler {
	return func(ctx services.Request) (v interface{}, err error) {
		v, err = next(ctx)
		return
	}
}
