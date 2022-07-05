package users

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/cluster/users/repository"
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
// @validate true
// @authorization false
// @permission false
// @internal false
// @title Get
// @description >>>
// Get a user
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | users_get_failed         | 500     | get user failed            |
// | users_get_nothing        | 404     | user was not found            |
// <<<
func get(ctx context.Context, argument GetArgument) (result *User, err errors.CodeError) {
	log := service.GetLog(ctx)
	row := repository.UserRow{}
	fetched, queryErr := postgres.QueryOne(
		ctx,
		postgres.NewConditions(postgres.Eq("ID", argument.Id)),
		&row,
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
	if !fetched {
		err = errors.NotFound("users_get_nothing").WithMeta("id", argument.Id)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = &User{
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
