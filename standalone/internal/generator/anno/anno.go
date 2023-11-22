package anno

import (
	"context"
	"github.com/aacfactory/gcg"
)

type ExampleAnno struct {
}

func (w *ExampleAnno) Annotation() (annotation string) {
	annotation = "example"
	return
}

func (w *ExampleAnno) HandleBefore(ctx context.Context, params []string, hasFnParam bool, hasFnResult bool) (code gcg.Code, err error) {
	return
}

func (w *ExampleAnno) HandleAfter(ctx context.Context, params []string, hasFnParam bool, hasFnResult bool) (code gcg.Code, err error) {
	return
}

func (w *ExampleAnno) ProxyBefore(ctx context.Context, params []string, hasFnParam bool, hasFnResult bool) (code gcg.Code, err error) {
	return
}

func (w *ExampleAnno) ProxyAfter(ctx context.Context, params []string, hasFnParam bool, hasFnResult bool) (code gcg.Code, err error) {
	return
}
