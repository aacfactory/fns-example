package users

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/cluster/users/repository"
	"github.com/aacfactory/fns/commons/uid"
	"github.com/aacfactory/fns/endpoints/authorizations"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/json"
	"time"
)

// CreateArgument
// @title Create user argument
// @description Create user argument
type CreateArgument struct {
	// Nickname
	// @title Nickname
	// @description Nickname
	Nickname string `json:"nickname" validate:"not_blank" message:"nickname is invalid"`
	// Mobile
	// @title Mobile
	// @description Mobile
	Mobile string `json:"mobile" validate:"not_blank" message:"mobile is invalid"`
	// Gender
	// @title Gender
	// @enum F(female),M(male),N(unknown)
	// @description Gender
	Gender string `json:"gender" validate:"oneof=F M N" message:"gender is invalid"`
	// Birthday
	// @title Birthday
	// @description Birthday
	Birthday json.Date `json:"birthday" validate:"required" message:"birthday is invalid"`
}

// CreateResult
// @title Create user result
// @description Create user result
type CreateResult struct {
	// Id
	// @title id
	// @description user id
	Id string `json:"id"`
	// Token
	// @title token
	// @description user token
	Token string `json:"token"`
}

// create
// @fn create
// @validate true
// @authorization false
// @permission false
// @internal false
// @transactional sql
// @title Create user
// @description >>>
// Create a user
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | users_create_failed      | 500     | create user failed            |
// <<<
func create(ctx context.Context, argument CreateArgument) (result *CreateResult, err errors.CodeError) {
	log := service.GetLog(ctx)
	userId := uid.UID()
	row := repository.UserRow{
		Id:       userId,
		CreateBY: userId,
		CreateAT: time.Now(),
		ModifyBY: "",
		ModifyAT: time.Time{},
		DeleteBY: "",
		DeleteAT: time.Time{},
		Version:  0,
		Nickname: argument.Nickname,
		Mobile:   argument.Mobile,
		Gender:   argument.Gender,
		Birthday: argument.Birthday.ToTime(),
		Avatar:   &repository.Avatar{},
	}
	insertErr := postgres.Insert(ctx, &row)
	if insertErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(insertErr).Message("users: create failed")
		}
		err = errors.ServiceError("users_create_failed").WithCause(insertErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}

	token, encodeErr := authorizations.Encode(ctx, userId, json.NewObject())
	if encodeErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(encodeErr).Message("users: create failed")
		}
		err = errors.ServiceError("users_create_failed").WithCause(encodeErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = &CreateResult{
		Id:    row.Id,
		Token: token,
	}
	return
}
