package users

// User
// @title 用户
// @description 用户
type User struct {
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
	// @description 性别
	Gender string `json:"gender"`
	// Age
	// @title 年龄
	// @description 年龄
	Age int `json:"age"`
	// Avatar
	// @title 头像图片地址
	// @description 头像图片地址
	Avatar string `json:"avatar"`
}
