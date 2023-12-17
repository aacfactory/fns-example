// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package modules

import (
	"github.com/aacfactory/fns-example/cluster/posts/modules/posts"
	"github.com/aacfactory/fns/services"
)

func endpoints() (v []services.Service) {
	v = []services.Service{
		posts.Service(),
	}
	return
}
