package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/copier"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/standalone/repository"
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
// @validate true
// @authorization false
// @permission false
// @internal false
// @title List
// @description >>>
// List posts
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | posts_list_failed        | 500     | list posts failed             |
// <<<
func list(ctx context.Context, argument ListArgument) (result []*Post, err errors.CodeError) {
	log := service.GetLog(ctx)
	rows := make([]*repository.PostRow, 0, 1)
	fetched, queryErr := postgres.QueryWithRange(
		ctx,
		postgres.NewConditions(postgres.NotEq("TITLE", "")),
		postgres.NewOrders().Desc("ID"),
		postgres.NewRange(argument.Offset, argument.Length),
		&rows,
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
	if !fetched {
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
