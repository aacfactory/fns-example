package components

import "github.com/aacfactory/fns/service"

// HelloComponent
// @components
type HelloComponent struct {
}

func (h *HelloComponent) Name() (name string) {
	name = "hello_component"
	return
}

func (h *HelloComponent) Build(options service.ComponentOptions) (err error) {
	return
}

func (h *HelloComponent) Close() {
	return
}
