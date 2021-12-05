package posts

import (
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-example/standalone/repository"
	"time"
)

type GetParam struct {
	Id string `json:"id,omitempty"`
}

func get(ctx fns.Context, param GetParam) (v *fns.Empty, err errors.CodeError) {

	x := &repository.PostRow{
		Id: param.Id,
	}

	dao := sql.DAO(ctx)

	has, getErr := dao.Get(ctx, x)
	fmt.Println(has, getErr)
	fmt.Println(fmt.Sprintf("%+v", *x))

	fmt.Println("Author", fmt.Sprintf("%+v", x.Author))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].User))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].Post))
	fmt.Println("Comments", fmt.Sprintf("%+v", x.Comments[0].Post.Comments[0]))

	x.Comments = make([]*repository.PostCommentRow, 0, 1)
	for i := 0; i < 2; i++ {
		x.Comments = append(x.Comments, &repository.PostCommentRow{
			Id:       fns.UID(),
			CreateBY: "-",
			CreateAT: time.Now(),
			Post:     x,
			User:     x.Author,
			Content:  fmt.Sprintf("%d", i),
		})
	}
	x.Title = "saved"

	saved, saveErr := dao.Save(ctx, x)
	fmt.Println(saved, saveErr)
	v = &fns.Empty{}
	fmt.Println(".............get..................")
	return
}
