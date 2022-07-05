package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/standalone/repository"
	"github.com/aacfactory/fns/service"
)

// CreateLikeArgument
// @title Create post like argument
// @description Create post like argument
type CreateLikeArgument struct {
	// PostId
	// @title post id
	// @description post id
	PostId string `json:"postId" validate:"required" message:"postId is invalid"`
}

// createLike
// @fn like
// @validate true
// @authorization true
// @permission false
// @internal false
// @transactional postgres
// @title Create post like
// @description >>>
// Create a post like
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | posts_like_failed        | 500     | create post like failed       |
// <<<
func createLike(ctx context.Context, argument CreateLikeArgument) (result *service.Empty, err errors.CodeError) {
	log := service.GetLog(ctx)
	request, hasRequest := service.GetRequest(ctx)
	if !hasRequest {
		err = errors.ServiceError("posts_like_failed").WithCause(fmt.Errorf("posts: bad request"))
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	userId := request.User().Id()
	row := repository.PostLikeRow{
		Id:     0,
		PostId: argument.PostId,
		UserId: userId,
	}
	insertErr := postgres.Insert(ctx, &row)
	if insertErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(insertErr).Message("posts: create comment failed")
		}
		err = errors.ServiceError("posts_like_failed").WithCause(insertErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = &service.Empty{}
	return
}
