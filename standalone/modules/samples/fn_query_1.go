package samples

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

// QueryParam
// @title 条形码解码表单
// @description 条形码解码表单
type QueryParam struct {
	// Title
	// @title Title
	// @description Title
	Title string `json:"title" validate:"required" message:"title is invalid"`
}

// queryBarcodeDecode
// @fn barcode_decode_query
// @validate false
// @authorization false
// @permission false
// @title 条形码解码
// @description 条形码解码
func queryBarcodeDecode(ctx fns.Context, param BarCodeDecodeParam) (v []*Barcode, err errors.CodeError) {

	return
}

type Barcodes []*Barcode

// queryBarcodeDecode2
// @fn barcode_decode_query_2
// @validate false
// @authorization false
// @permission false
// @title 条形码解码
// @description 条形码解码
func queryBarcodeDecode2(ctx fns.Context, param BarCodeDecodeParam) (v Barcodes, err errors.CodeError) {

	return
}
