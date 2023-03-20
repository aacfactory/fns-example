package users

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-contrib/databases/sql/dal"
	"github.com/aacfactory/fns-example/standalone/repositories/postgres"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/json"
)

// GetArgument
// @title Get user argument
// @description Get user argument
type GetArgument struct {
	// Id
	// @title Id
	// @description Id
	Id string `json:"id" validate:"required" message:"id is invalid"`
}

// get
// @fn get
// @validation
// @errors >>>
// + users_get_failed
//   - en: users get failed
//
// + users_get_nothing
//   - en: users get nothing
//
// <<<
// @title Get
// @description >>>
// Get a user
// <<<
func get(ctx context.Context, argument GetArgument) (result User, err errors.CodeError) {
	log := service.GetLog(ctx)
	row, queryErr := dal.QueryOne[*postgres.UserRow](
		ctx,
		dal.NewConditions(dal.Eq("ID", argument.Id)).And(dal.Eq("BD", sql.NewDate(2022, 1, 1))).And(dal.Eq("BT", sql.NewTime(10, 12, 59))),
	)
	if queryErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(queryErr).Message("users: get failed")
		}
		err = errors.ServiceError("users_get_failed").WithCause(queryErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	if row == nil {
		err = errors.NotFound("users_get_nothing").WithMeta("id", argument.Id)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	log.Debug().Message(fmt.Sprintf("--------%s---%v---%v", row.Id, row.BD, row.BT))
	result = User{
		Id:       row.Id,
		CreateAT: row.CreateAT,
		Nickname: row.Nickname,
		Mobile:   row.Mobile,
		Gender:   row.Gender,
		Birthday: json.NewDateFromTime(row.Birthday),
		Avatar:   row.Avatar,
	}
	return
}
