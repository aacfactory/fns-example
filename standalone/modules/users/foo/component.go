package foo

import "github.com/aacfactory/fns/service"

type Foo struct {
}

func (c *Foo) Name() (name string) {
	name = "c"
	return
}

func (c *Foo) Build(options service.ComponentOptions) (err error) {
	return
}

func (c *Foo) Close() {
	return
}
