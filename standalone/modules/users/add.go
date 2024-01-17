package users

import (
	"github.com/aacfactory/fns/commons/times"
	"github.com/aacfactory/fns/context"
)

// AddParam
// @title add param
// @description add user param
type AddParam struct {
	// Id
	// @title user id
	// @description user id
	Id string `json:"id" validate:"required" validate-message:""`
	// Name
	// @title name
	// @description name
	Name string `json:"name" validate:"required" validate-message:""`
	// Age
	// @title age
	// @description age
	Age int `json:"age"`
	// Birthday
	// @title birthday
	// @description birthday
	Birthday times.Time `json:"birthday"`
}

// add
// @fn add
// @authorization
// @validation
// @cache set
// @title add
// @description >>>
// add user
// <<<
// @errors >>>
// user_not_found
// zh: zh_message
// en: en_message
// <<<
func add(ctx context.Context, param AddParam) (v User, err error) {

	return
}
