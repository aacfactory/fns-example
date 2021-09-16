package posts

import (
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-example/standalone/repository"
	"time"
)

type CreateParam struct {
	Title   string `json:"title" validate:"required" message:"title is invalid"`
	Content string `json:"content" validate:"not_blank" message:"content is invalid"`
}

// create
// @fn create
// @validate true
// @sqlTX true
// @authorization true
// @permission true
// @description foo
func create(ctx fns.Context, param CreateParam) (row0 *repository.PostRow, err errors.CodeError) {

	row := &repository.PostRow{
		Id:       fns.UID(),
		CreateBY: "-",
		CreateAT: time.Time{},
		Version:  1,
		Title:    param.Title,
		Content:  param.Content,
		Author: &repository.UserRow{
			Id: "1",
		},
		Comments: nil,
	}

	comments := make([]*repository.PostCommentRow, 0, 1)
	comments = append(comments, &repository.PostCommentRow{
		Id:       fns.UID(),
		CreateBY: "-",
		CreateAT: time.Time{},
		Post:     row,
		User:     row.Author,
		Content:  "foo",
	})

	row.Comments = comments

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

	x := &repository.PostRow{
		Id: row.Id,
	}

	has, getErr := sql.DAO(x).Get(ctx)
	fmt.Println(has, getErr)
	fmt.Println(fmt.Sprintf("%+v", *x))

	fmt.Println("Author", fmt.Sprintf("%+v", x.Author))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].User))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].Post))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].Post.Comments[0]))
	return
}
