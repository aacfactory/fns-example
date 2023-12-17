package posts

import (
	"github.com/aacfactory/copier"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/cluster/posts/repositories"
	"github.com/aacfactory/fns/context"
)

type ListParam struct {
	UserId string `json:"userId" validate:"required" validate-message:"user_id_required"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
}

// list
// @fn list
// @readonly
// @validation
// @cache-control max-age=10 public=true
// @barrier
// @metric
// @title list
// @description >>>
// list posts
// <<<
// @errors >>>
// posts_not_found
// zh: zh_message
// en: en_message
// <<<
func list(ctx context.Context, param ListParam) (v Posts, err error) {
	rows, queryErr := postgres.Query[repositories.PostRow](
		ctx,
		param.Offset, param.Length,
		postgres.Conditions(postgres.Eq("UserId", param.UserId)),
	)
	if queryErr != nil {
		err = errors.ServiceError("posts: list failed").WithCause(queryErr)
		return
	}
	rowsLen := len(rows)
	if rowsLen == 0 {
		return
	}
	v = make(Posts, 0, rowsLen)
	if err = copier.Copy(&v, rows); err != nil {
		err = errors.ServiceError("posts: list failed").WithCause(err)
		return
	}
	return
}
