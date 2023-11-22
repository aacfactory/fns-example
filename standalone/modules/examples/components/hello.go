package components

import (
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/services"
)

// HelloComponent
// @components
type HelloComponent struct {
}

func (component *HelloComponent) Name() (name string) {
	return "hello"
}

func (component *HelloComponent) Construct(options services.Options) (err error) {
	return
}

func (component *HelloComponent) Shutdown(ctx context.Context) {
}
