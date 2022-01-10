package samples

import (
	"github.com/aacfactory/fns-example/standalone/modules/users"
	"github.com/aacfactory/json"
)

// Sample
// @title title of sample
// @description >>> description of sample
// > support markdown
//
// <<<
type Sample struct {
	// Id
	// @title 编号
	// @description 编号
	Id string `json:"id"`
	// Mobile
	// @title 手机号
	// @description 手机号
	Mobile string `json:"mobile"`
	// Name
	// @title 姓名
	// @description 姓名
	Name string `json:"name"`
	// Gender
	// @title 性别
	// @enum M,F,N
	// @description 性别
	Gender string `json:"gender"`
	// Age
	// @title 年龄
	// @description 年龄
	// @enum 1,2,3
	Age int `json:"age"`
	// Avatar
	// @title 头像图片地址
	// @description 头像图片地址
	Avatar string `json:"avatar"`
	// Score
	// @title Score
	// @description Score
	Score float32 `json:"score,omitempty"`
	// DOB
	// @title DOB
	// @description DOB
	DOB json.Date `json:"dob,omitempty"`
	// CreateAT
	// @title CreateAT
	// @description CreateAT
	CreateAT json.Time `json:"createAt,omitempty"`
	// Tokens
	// @title Tokens
	// @description Tokens
	Tokens []string
	// Users
	// @title Users
	// @description Users
	Users []*users.User `json:"users,omitempty"`
	// UserMap
	// @title UserMap
	// @description UserMap
	UserMap map[string]*users.User `json:"userMap,omitempty"`
	// Raw
	// @title Raw
	// @description Raw
	Raw json.RawMessage
}
