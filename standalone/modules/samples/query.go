package samples

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

// QueryParam
// @title title
// @description description
type QueryParam struct {
	// Offset
	// @title title
	// @description description
	Offset int `json:"offset" validate:"required" message:"offset is invalid"`
	// Limit
	// @title title
	// @description description
	Limit int `json:"limit" validate:"required" message:"limit is invalid"`
}

// query
// @fn query
// @validate true
// @authorization true
// @permission false
// @title query
// @description query
func query(ctx fns.Context, param QueryParam) (v []*Sample, err errors.CodeError) {

	return
}
