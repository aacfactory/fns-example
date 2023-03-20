package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/copier"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/sql/dal"
	"github.com/aacfactory/fns-example/standalone/repositories/postgres"
	"github.com/aacfactory/fns/service"
)

// ListArgument
// @title List posts argument
// @description List argument
type ListArgument struct {
	Offset int `json:"offset" validate:"min=0" message:"offset is invalid"`
	Length int `json:"length" validate:"min=10,max=50" message:"length is invalid"`
}

// list
// @fn list
// @validation
// @barrier
// @errors >>>
// + posts_list_failed
//   - en: posts list failed
//
// <<<
// @title List
// @description >>>
// List posts
// <<<
func list(ctx context.Context, argument ListArgument) (result Posts, err errors.CodeError) {
	log := service.GetLog(ctx)
	rows, queryErr := dal.QueryWithRange[*postgres.PostRow](
		ctx,
		dal.NewConditions(dal.NotEq("TITLE", "")),
		dal.NewOrders().Desc("ID"),
		dal.NewRange(argument.Offset, argument.Length),
	)
	if queryErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(queryErr).Message("posts: list failed")
		}
		err = errors.ServiceError("posts_list_failed").WithCause(queryErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = Posts{}
	if rows == nil || len(rows) == 0 {
		return
	}
	result = make([]*Post, 0, len(rows))
	cpErr := copier.Copy(&result, &rows)
	if cpErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(cpErr).Message("posts: list failed")
		}
		err = errors.ServiceError("posts_list_failed").WithCause(cpErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	return
}
