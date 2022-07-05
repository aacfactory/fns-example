package posts

import (
	"context"
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-example/standalone/repository"
	"github.com/aacfactory/fns/commons/uid"
	"github.com/aacfactory/fns/service"
	"time"
)

// CreateArgument
// @title Create post argument
// @description Create post argument
type CreateArgument struct {
	// Title
	// @title post title
	// @description post title
	Title string `json:"title" validate:"required" message:"title is invalid"`
	// Content
	// @title post content
	// @description post content
	Content string `json:"content" validate:"required" message:"content is invalid"`
}

// CreateResult
// @title Create post result
// @description Create post result
type CreateResult struct {
	// Id
	// @title post id
	// @description post id
	Id string `json:"id"`
}

// create
// @fn create
// @validate true
// @authorization true
// @permission false
// @internal false
// @transactional postgres
// @title Create post
// @description >>>
// Create a post
// ----------
// errors:
// | Name                     | Code    | Description                   |
// |--------------------------|---------|-------------------------------|
// | posts_create_failed      | 500     | create post failed            |
// <<<
func create(ctx context.Context, argument CreateArgument) (result *CreateResult, err errors.CodeError) {
	log := service.GetLog(ctx)
	request, hasRequest := service.GetRequest(ctx)
	if !hasRequest {
		err = errors.ServiceError("posts_create_failed").WithCause(fmt.Errorf("posts: bad request"))
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	userId := request.User().Id()
	row := repository.PostRow{
		Id: uid.UID(),
		User: &repository.UserRow{
			Id: userId,
		},
		CreateAT: time.Now(),
		Version:  0,
		Title:    argument.Title,
		Content:  argument.Content,
		Comments: nil,
		Likes:    0,
	}
	insertErr := postgres.Insert(ctx, &row)
	if insertErr != nil {
		if log.ErrorEnabled() {
			log.Error().Caller().Cause(insertErr).Message("posts: create failed")
		}
		err = errors.ServiceError("posts_create_failed").WithCause(insertErr)
		if log.DebugEnabled() {
			log.Debug().Caller().Message(fmt.Sprintf("%+v", err))
		}
		return
	}
	result = &CreateResult{
		Id: row.Id,
	}
	return
}
