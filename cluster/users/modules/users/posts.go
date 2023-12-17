package users

import (
	"github.com/aacfactory/fns-example/cluster/posts/modules/posts"
	"github.com/aacfactory/fns/context"
)

type PostsParam struct {
	UserId string `json:"userId" validate:"required" validate-message:"user_id_required"`
}

// post
// @fn post
// @readonly
// @validation
// @cache-control max-age=10 public=true
// @barrier
// @metric
// @title post
// @description >>>
// list user posts
// <<<
// @errors >>>
// posts_not_found
// zh: zh_message
// en: en_message
// <<<
func post(ctx context.Context, param PostsParam) (v posts.Posts, err error) {
	v, err = posts.List(ctx, posts.ListParam{
		UserId: param.UserId,
		Offset: 0,
		Length: 10,
	})
	return
}
