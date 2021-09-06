package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns/secret"
)

type LoginParam struct {
	Id       string `json:"id,omitempty" validate:"required" message:"id is invalid"`
	Password string `json:"password,omitempty" validate:"required" message:"password is invalid"`
}

type Token struct {
	Value string `json:"value,omitempty"`
}

func login(ctx fns.Context, param LoginParam) (token *Token, err errors.CodeError) {
	user, getErr := get(ctx, GetParam{
		Id: param.Id,
	})

	if getErr != nil {
		err = errors.BadRequest("user was not found")
		return
	}

	if !secret.ValidatePassword([]byte(user.Password), []byte(param.Password)) {
		err = errors.BadRequest("invalid password")
		return
	}

	_ = ctx.User().Attributes().Put("id", user.Id)
	_ = ctx.User().Attributes().Put("gender", user.Gender)

	value, encodeErr := ctx.App().Authorizations().Encode(ctx)
	if encodeErr != nil {
		err = errors.ServiceError("encode failed").WithCause(encodeErr)
		return
	}

	token = &Token{
		Value: string(value),
	}
	return
}
