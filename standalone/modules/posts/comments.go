package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/standalone/repository"
	"github.com/aacfactory/fns/service"
	"time"
)

// CreateCommentArgument
// @title Create post comment argument
// @description Create post comment argument
type CreateCommentArgument struct {
	// PostId
	// @title post id
	// @description post id
	PostId string `json:"postId" validate:"required" message:"postId is invalid"`
	// Content
	// @title content
	// @description content
	Content string `json:"content" validate:"required" message:"content is invalid"`
}

// CreateCommentResult
// @title Create post comment result
// @description Create post comment result
type CreateCommentResult struct {
	// Id
	// @title post comment id
	// @description post comment id
	Id int64 `json:"id"`
}

// createComment
// @fn create_comment
// @validate true
// @authorization true
// @permission false
// @internal false
// @transactional sql
// @title Create post comment
// @description >>>
// Create a post comment
// ----------
// errors:
// | Name                           | Code    | Description                   |
// |--------------------------------|---------|-------------------------------|
// | posts_create_comment_failed    | 500     | create post comment failed    |
// <<<
func createComment(ctx context.Context, argument CreateCommentArgument) (result *CreateCommentResult, err errors.CodeError) {
	log := service.GetLog(ctx)
	request, hasRequest := service.GetRequest(ctx)
	if !hasRequest {
		err = errors.ServiceError("posts_create_comment_failed").WithCause(fmt.Errorf("posts: bad request"))
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	userId := request.User().Id()
	row := repository.PostCommentRow{
		Id:     0,
		PostId: argument.PostId,
		User: &repository.UserRow{
			Id: userId,
		},
		CreateAT: time.Now(),
		Content:  argument.Content,
	}
	insertErr := postgres.Insert(ctx, &row)
	if insertErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(insertErr).Message("posts: create comment failed")
		}
		err = errors.ServiceError("posts_create_comment_failed").WithCause(insertErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = &CreateCommentResult{
		Id: row.Id,
	}
	return
}
