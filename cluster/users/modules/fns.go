// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package modules

import (
	"github.com/aacfactory/fns-example/cluster/users/modules/users"
	"github.com/aacfactory/fns/services"
)

func endpoints() (v []services.Service) {
	v = []services.Service{
		users.Service(),
	}
	return
}
