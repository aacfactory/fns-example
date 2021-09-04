package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
)

type ListParam struct {
	Offset int `json:"offset,omitempty" validate:"min=0,max=999" message:"offset is invalid"`
	Length int `json:"length,omitempty" validate:"min=10,max=30" message:"length is invalid"`
}

const (
	listSQL = `SELECT * FROM "FNS"."USER" OFFSET $1 LIMIT $2`
)

func list(ctx fns.Context, param ListParam) (users []*User, err errors.CodeError) {

	offset := param.Offset
	length := param.Length

	tuple := sql.NewTuple()
	tupleErr := tuple.Append(offset, length)
	if tupleErr != nil {
		ctx.App().Log().Error().Caller().Cause(tupleErr).Message("append sql arg failed")
		err = errors.ServiceError("app sql arg failed").WithCause(tupleErr)
		return
	}

	rows, queryErr := sql.Query(ctx, sql.Param{
		Query: listSQL,
		Args:  tuple,
		InTx:  false,
	})

	if queryErr != nil {
		ctx.App().Log().Error().Caller().Cause(queryErr).Message("query failed")
		err = errors.ServiceError("query failed").WithCause(queryErr)
		return
	}

	if rows.Empty() {
		return
	}

	userRows := make([]*UserRow, 0, 1)
	scanErr := rows.Scan(&userRows)
	if scanErr != nil {
		ctx.App().Log().Error().Caller().Cause(scanErr).Message("scan failed")
		err = errors.ServiceError("scan failed").WithCause(scanErr)
		return
	}

	for _, row := range userRows {
		users = append(users, &User{
			Id:         row.Id,
			Name:       row.Name,
			Password:   row.Password,
			Gender:     row.Gender,
			Age:        row.Age,
			Active:     row.Active,
			SignUpTime: row.SignUpTime,
			Profile:    row.Profile,
			Score:      row.Score,
			DOB:        row.DOB,
		})
	}

	if ctx.App().Log().DebugEnabled() {
		ctx.App().Log().Debug().Message("user list succeed")
	}

	return
}
