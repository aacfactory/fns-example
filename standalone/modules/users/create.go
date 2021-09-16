package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-example/standalone/repository"
	"github.com/aacfactory/fns/secret"
	"github.com/aacfactory/json"
	"time"
)

type CreateParam struct {
	Name     string    `json:"name,omitempty" validate:"required" message:"name is invalid"`
	Password string    `json:"password,omitempty" validate:"required" message:"password is invalid"`
	Gender   string    `json:"gender,omitempty" validate:"required" message:"gender is invalid"`
	Age      int       `json:"age,omitempty" validate:"min=1,max=100" message:"age is invalid"`
	Active   bool      `json:"active,omitempty" validate:"required" message:"active is invalid"`
	Score    float64   `json:"score,omitempty" validate:"required" message:"score is invalid"`
	DOB      json.Date `json:"dob,omitempty" validate:"required" message:"dob is invalid"`
}

// create
// @fn create
// @validate true
// @sqlTX true
// @authorization true
// @permission true
// @description foo
func create(ctx fns.Context, param CreateParam) (err errors.CodeError) {

	password, _ := secret.HashPassword([]byte(param.Password))

	row := &repository.UserRow{
		Id:         fns.UID(),
		CreateBY:   "-",
		CreateAT:   time.Now(),
		Version:    1,
		Name:       param.Name,
		Password:   string(password),
		Gender:     param.Gender,
		Age:        param.Age,
		Active:     true,
		SignUpTime: time.Now(),
		Profile: &repository.UserProfile{
			Name: param.Name,
			Age:  param.Age,
		},
		Score: param.Score,
		DOB:   param.DOB.ToTime(),
	}

	affected, insertErr := sql.DAO(row).Insert(ctx)

	if insertErr != nil {
		ctx.App().Log().Error().Caller().Cause(insertErr).Message("execute failed")
		err = errors.ServiceError("execute failed").WithCause(insertErr)
		return
	}

	if affected < 1 {
		ctx.App().Log().Error().Caller().Message("execute failed no affected")
		err = errors.ServiceError("execute failed for no affected")
		return
	}

	return
}
