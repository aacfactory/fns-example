package samples

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

// SettingsQueryParam
// @title 获取系统设置表单
// @description 获取系统设置表单
type SettingsQueryParam struct {
	// Keys
	// @title 键值列表
	// @description 键值列表
	Keys []string `json:"keys"`
}

// settingsQuery
// @fn settings_query
// @validate false
// @authorization false
// @permission false
// @title 获取系统设置
// @description 获取系统设置
func settingsQuery(ctx fns.Context, param SettingsQueryParam) (views []*Settings, err errors.CodeError) {

	return
}

