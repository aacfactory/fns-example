package users

import (
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

// get
// @fn get
// @readonly
// @authorization
// @permission
// @validation
// @cache get-set 5
// @cache-control max-age=10 public=true
// @barrier
// @metric
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
	return
}
