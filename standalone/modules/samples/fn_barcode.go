package samples

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

// BarCodeDecodeParam
// @title 条形码解码表单
// @description 条形码解码表单
type BarCodeDecodeParam struct {
	// Code
	// @title 条形码类型
	// @description 条形码类型 code_128 upc_a upc_e
	Kind string `json:"kind,omitempty"`
	// Data
	// @title BASE64 图片
	// @description BASE64 图片
	Data string `json:"data" validate:"not_blank" message:"图片数据不能为空"`
}

// Barcode
// @title 条形码
// @description 条形码
type Barcode struct {
	// Code
	// @title 条形码
	// @description 条形码
	Code string `json:"code,omitempty"`
}

// barcodeDecode
// @fn barcode_decode
// @validate false
// @authorization false
// @permission false
// @title 条形码解码
// @description 条形码解码
func barcodeDecode(ctx fns.Context, param BarCodeDecodeParam) (v *Barcode, err errors.CodeError) {

	return
}
