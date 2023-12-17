package users

import (
	"fmt"
	"github.com/aacfactory/fns-example/standalone/modules/examples/components"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/commons"
	"time"
)

type Users []User

// list
// @fn list
// @readonly
// @authorization
// @permission
// @validation
// @cache get-set 5
// @cache-control max-age=10 public=true
// @barrier
// @metric
// @middlewares >>>
// Middle
// github.com/aacfactory/fns-example/standalone/modules/users/middles.Middle
// <<<
// @title list
// @description >>>
// dafasdf
// adsfasfd
// <<<
// @errors >>>
// user_not_found
// zh: zh_message
// en: en_message
// <<<
func list(ctx context.Context) (v Users, err error) {
	v = []User{User{
		Id:       "1",
		Name:     "1",
		Age:      "1",
		Birthday: time.Now(),
	}}
	c, _ := Component[*components.HelloComponent](ctx, "hello")
	c.Name()
	return
}

type Middle struct {
}

func (m Middle) Handler(next commons.FnHandler) commons.FnHandler {
	return func(r services.Request) (v interface{}, err error) {
		fmt.Println("1")
		v, err = next(r)
		return
	}
}
