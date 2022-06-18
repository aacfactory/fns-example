package users

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-example/standalone/modules/users/foo"
	"github.com/aacfactory/fns/service"
)

// fooLoader
// @component foo:github.com/aacfactory/fns-example/standalone/modules/users/foo.Foo
func fooLoader() service.Component {
	return &foo.Foo{}
}

// c1Loader
// @component c:c1
func c1Loader() service.Component {
	return &c1{}
}

type c1 struct {
}

func (c *c1) Name() (name string) {
	name = "c"
	return
}

func (c *c1) Build(options service.ComponentOptions) (err error) {
	return
}

func (c *c1) Close() {
	return
}

func getComponentC2(ctx context.Context) (v *c1) {
	c, has := service.GetComponent(ctx, "")
	if !has {
		panic(fmt.Sprintf("%+v", errors.Warning("%s: get c1 component failed cause not found in context")))
		return
	}
	ok := false
	v, ok = c.(*c1)
	if !ok {
		panic(fmt.Sprintf("%+v", errors.Warning("%s: get c1 component failed cause type is not matched")))
		return
	}
	return
}
