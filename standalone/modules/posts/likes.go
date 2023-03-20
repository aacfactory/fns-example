package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/sql/dal"
	"github.com/aacfactory/fns-example/standalone/repositories/postgres"
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
// @validation
// @authorization
// @errors >>>
// + posts_like_failed
//   - en: posts like failed
//
// <<<
// @sql postgres
// @transactional
// @title Create post like
// @description >>>
// Create a post like
// <<<
func createLike(ctx context.Context, argument CreateLikeArgument) (result service.Empty, err errors.CodeError) {
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
	row := postgres.PostLikeRow{
		Id:     0,
		PostId: argument.PostId,
		UserId: userId.String(),
	}
	insertErr := dal.Insert(ctx, &row)
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
	return
}
