package middles

import (
	"fmt"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/commons"
)

type Middle struct {
}

func (m Middle) Handler(next commons.FnHandler) commons.FnHandler {
	return func(r services.Request) (v interface{}, err error) {
		fmt.Println("1")
		v, err = next(r)
		return
	}
}
