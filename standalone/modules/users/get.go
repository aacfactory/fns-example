package users

import (
	"fmt"
	"github.com/aacfactory/fns-example/standalone/modules/examples/components"
	"github.com/aacfactory/fns/context"
	"time"
)

// GetParam
// @title get param
// @description get param
type GetParam struct {
	// Id
	// @title id
	// @description id
	// @validate-message-i18n >>>
	// zh: zh_message
	// en: en_message
	// <<<
	Id string `json:"id" validate:"not_blank" validate-message:"invalid_id"`
}

func (param GetParam) CacheKey(ctx context.Context) (key []byte, err error) {
	key = []byte(fmt.Sprintf("users:%s", param.Id))
	return
}

// get
// @fn get
// @readonly
// @validation
// @metric
// @cache get-set 50
// @cache-control max-age=100 public=true
// @barrier
// @title get
// @description >>>
// dafasdf
// adsfasfd
// <<<
// @errors >>>
// user_not_found
// zh: zh_message
// en: en_message
// <<<
func get(ctx context.Context, param GetParam) (v User, err error) {
	v = User{
		Id:       "1",
		Name:     "1",
		Age:      "1",
		Birthday: time.Now(),
	}
	c, _ := Component[*components.HelloComponent](ctx, "hello")
	c.Name()
	return
}
