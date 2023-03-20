// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package modules

import (
	"github.com/aacfactory/fns-example/standalone/modules/examples"
	"github.com/aacfactory/fns-example/standalone/modules/posts"
	"github.com/aacfactory/fns-example/standalone/modules/users"
	"github.com/aacfactory/fns/service"
)

func services() (v []service.Service) {
	v = []service.Service{
		examples.Service(),
		posts.Service(),
		users.Service(),
	}
	return
}
