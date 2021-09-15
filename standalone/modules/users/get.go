package users

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
)

type GetParam struct {
	Id string `json:"id,omitempty" validate:"required" message:"id is invalid"`
}

func get(ctx fns.Context, param GetParam) (user *User, err errors.CodeError) {
	const (
		_query = `SELECT * FROM "FNS"."USER" WHERE "ID" = $1`
	)

	id := param.Id

	tuple := sql.NewTuple()
	tuple.Append(id)

	rows, queryErr := sql.Query(ctx, sql.Param{
		Query: _query,
		Args:  tuple,
	})

	if queryErr != nil {
		ctx.App().Log().Error().Caller().Cause(queryErr).Message("query failed")
		err = errors.ServiceError("query failed").WithCause(queryErr)
		return
	}

	if rows.Empty() {
		return
	}

	userRow := &UserRow{}
	scanErr := rows.Scan(userRow)
	if scanErr != nil {
		ctx.App().Log().Error().Caller().Cause(scanErr).Message("scan failed")
		err = errors.ServiceError("scan failed").WithCause(scanErr)
		return
	}

	user = UserViewMapFromTableRow(userRow)

	return
}
