package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns/secret"
	"github.com/aacfactory/json"
	"time"
)

type CreateParam struct {
	Id       string    `json:"id,omitempty" validate:"required" message:"id is invalid"`
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
// @authorization true
// @permission true
// @description foo
func create(ctx fns.Context, param CreateParam) (err errors.CodeError) {
	const (
		_query = `INSERT INTO "FNS"."USER" 
				  	(
				  	"ID", "NAME", "PASSWORD", "AGE", "ACTIVE", "SIGN_UP_TIME", "PROFILE", "SCORE", "DOB", "GENDER"
					)
				  	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	)
	password, _ := secret.HashPassword([]byte(param.Password))
	signupTime := time.Now()

	//dob, _ := time.Parse("2006-01-02", param.DOB)

	tuple := sql.NewTuple()
	tuple.Append(param.Id, param.Name, string(password), param.Age, true, signupTime, []byte("{}"), param.Score, param.DOB, param.Gender)


	r, execErr := sql.Execute(ctx, sql.Param{
		Query: _query,
		Args:  tuple,
		InTx:  false,
	})

	if execErr != nil {
		ctx.App().Log().Error().Caller().Cause(execErr).Message("execute failed")
		err = errors.ServiceError("execute failed").WithCause(execErr)
		return
	}

	if r.Affected < 1 {
		ctx.App().Log().Error().Caller().Cause(execErr).Message("execute failed no affected")
		err = errors.ServiceError("execute failed for no affected").WithCause(execErr)
		return
	}

	return
}
