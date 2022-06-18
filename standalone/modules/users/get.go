package users

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
)

// GetArgument
// @title 获取参数
// @description 获取参数
type GetArgument struct {
	// Id
	// @title 藏品id
	// @description 藏品id
	Id int64 `json:"id" validate:"required" message:"id is invalid"`
}

// get
// @fn get
// @validate true
// @authorization false
// @permission false
// @internal false
// @title 获取用户信息
// @description >>>
// 获取用户信息
// ----------
// errors:
// * user_get_failed
// <<<
func get(ctx context.Context, argument GetArgument) (v *User, err errors.CodeError) {
	v = &User{
		Id:     fmt.Sprintf("%v", argument.Id),
		Mobile: "000",
		Name:   "foo",
		Gender: "F",
		Age:    10,
		Avatar: "bar",
	}
	err = errors.ServiceError("").WithMeta("", "")

	return
}
