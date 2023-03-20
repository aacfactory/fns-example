// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package posts

import (
	"context"

	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns/endpoints/authorizations"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/documents"
	"github.com/aacfactory/fns/service/validators"
)

const (
	_name            = "posts"
	_createFn        = "create"
	_createCommentFn = "create_comment"
	_createLikeFn    = "like"
	_listFn          = "list"
)

func Create(ctx context.Context, argument CreateArgument) (result CreateArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("posts: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _createFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("posts: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _createFn).
			WithCause(scanErr)
		return
	}
	return
}

func CreateComment(ctx context.Context, argument CreateCommentArgument) (result CreateCommentArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("posts: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _createCommentFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("posts: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _createCommentFn).
			WithCause(scanErr)
		return
	}
	return
}

func CreateLike(ctx context.Context, argument CreateLikeArgument) (result service.Empty, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("posts: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _createLikeFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("posts: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _createLikeFn).
			WithCause(scanErr)
		return
	}
	return
}

func List(ctx context.Context, argument ListArgument) (result ListArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("posts: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _listFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("posts: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _listFn).
			WithCause(scanErr)
		return
	}
	return
}

func Service() (v service.Service) {
	v = &_service_{
		Abstract: service.NewAbstract(
			_name,
			false,
		),
	}
	return
}

type _service_ struct {
	service.Abstract
}

func (svc *_service_) Handle(ctx context.Context, fn string, argument service.Argument) (v interface{}, err errors.CodeError) {
	switch fn {
	case _createFn:
		// param
		param := CreateArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("posts: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// use sql database
		ctx = sql.WithOptions(ctx, sql.Database("postgres"))
		// sql begin transaction
		beginTransactionErr := sql.BeginTransaction(ctx)
		if beginTransactionErr != nil {
			err = errors.Warning("posts: begin sql transaction failed").WithCause(beginTransactionErr)
			return
		}
		// execute function
		v, err = create(ctx, param)
		// sql commit transaction
		if err == nil {
			commitTransactionErr := sql.CommitTransaction(ctx)
			if commitTransactionErr != nil {
				_ = sql.RollbackTransaction(ctx)
				err = errors.ServiceError("posts: commit sql transaction failed").WithCause(commitTransactionErr)
				return
			}
		}
		break
	case _createCommentFn:
		// verify authorizations
		err = authorizations.Verify(ctx)
		if err != nil {
			break
		}
		// param
		param := CreateCommentArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("posts: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// use sql database
		ctx = sql.WithOptions(ctx, sql.Database("postgres"))
		// sql begin transaction
		beginTransactionErr := sql.BeginTransaction(ctx)
		if beginTransactionErr != nil {
			err = errors.Warning("posts: begin sql transaction failed").WithCause(beginTransactionErr)
			return
		}
		// execute function
		v, err = createComment(ctx, param)
		// sql commit transaction
		if err == nil {
			commitTransactionErr := sql.CommitTransaction(ctx)
			if commitTransactionErr != nil {
				_ = sql.RollbackTransaction(ctx)
				err = errors.ServiceError("posts: commit sql transaction failed").WithCause(commitTransactionErr)
				return
			}
		}
		break
	case _createLikeFn:
		// verify authorizations
		err = authorizations.Verify(ctx)
		if err != nil {
			break
		}
		// param
		param := CreateLikeArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("posts: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// use sql database
		ctx = sql.WithOptions(ctx, sql.Database("postgres"))
		// sql begin transaction
		beginTransactionErr := sql.BeginTransaction(ctx)
		if beginTransactionErr != nil {
			err = errors.Warning("posts: begin sql transaction failed").WithCause(beginTransactionErr)
			return
		}
		// execute function
		v, err = createLike(ctx, param)
		// sql commit transaction
		if err == nil {
			commitTransactionErr := sql.CommitTransaction(ctx)
			if commitTransactionErr != nil {
				_ = sql.RollbackTransaction(ctx)
				err = errors.ServiceError("posts: commit sql transaction failed").WithCause(commitTransactionErr)
				return
			}
		}
		break
	case _listFn:
		// param
		param := ListArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("posts: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// barrier
		v, err = svc.Barrier(ctx, _listFn, argument, func() (v interface{}, err errors.CodeError) {
			// execute function
			v, err = list(ctx, param)
			return
		})
		break
	default:
		err = errors.Warning("posts: fn was not found").WithMeta("service", _name).WithMeta("fn", fn)
		break
	}
	return
}

func (svc *_service_) Document() (doc service.Document) {
	document := documents.NewService(_name, "Post service")
	// create
	document.AddFn(
		"create", "Create post", "Create a post", false, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "CreateArgument").
			SetTitle("Create post argument").
			SetDescription("Create post argument").
			AddProperty(
				"title",
				documents.String().
					SetTitle("post title").
					SetDescription("post title").
					AsRequired().
					SetValidation(documents.NewElementValidation("title is invalid")),
			).
			AddProperty(
				"content",
				documents.String().
					SetTitle("post content").
					SetDescription("post content").
					AsRequired().
					SetValidation(documents.NewElementValidation("content is invalid")),
			),
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "CreateResult").
			SetTitle("Create post result").
			SetDescription("Create post result").
			AddProperty(
				"id",
				documents.String().
					SetTitle("post id").
					SetDescription("post id"),
			),
		[]documents.FnError{
			{
				Name_: "posts_create_failed",
				Descriptions_: map[string]string{
					"en": "posts create failed",
				},
			},
		},
	)

	// create_comment
	document.AddFn(
		"create_comment", "Create post comment", "Create a post comment", true, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "CreateCommentArgument").
			SetTitle("Create post comment argument").
			SetDescription("Create post comment argument").
			AddProperty(
				"postId",
				documents.String().
					SetTitle("post id").
					SetDescription("post id").
					AsRequired().
					SetValidation(documents.NewElementValidation("postId is invalid")),
			).
			AddProperty(
				"content",
				documents.String().
					SetTitle("content").
					SetDescription("content").
					AsRequired().
					SetValidation(documents.NewElementValidation("content is invalid")),
			),
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "CreateCommentResult").
			SetTitle("Create post comment result").
			SetDescription("Create post comment result").
			AddProperty(
				"id",
				documents.Int64().
					SetTitle("post comment id").
					SetDescription("post comment id"),
			),
		[]documents.FnError{
			{
				Name_: "posts_create_comment_failed",
				Descriptions_: map[string]string{
					"en": "posts create comment failed",
				},
			},
		},
	)

	// like
	document.AddFn(
		"like", "Create post like", "Create a post like", true, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "CreateLikeArgument").
			SetTitle("Create post like argument").
			SetDescription("Create post like argument").
			AddProperty(
				"postId",
				documents.String().
					SetTitle("post id").
					SetDescription("post id").
					AsRequired().
					SetValidation(documents.NewElementValidation("postId is invalid")),
			),
		documents.Struct("github.com/aacfactory/fns/service", "Empty").
			SetDescription("Empty object"),
		[]documents.FnError{
			{
				Name_: "posts_like_failed",
				Descriptions_: map[string]string{
					"en": "posts like failed",
				},
			},
		},
	)

	// list
	document.AddFn(
		"list", "List", "List posts", false, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "ListArgument").
			SetTitle("List posts argument").
			SetDescription("List argument").
			AddProperty(
				"offset",
				documents.Int64().
					SetValidation(documents.NewElementValidation("offset is invalid")),
			).
			AddProperty(
				"length",
				documents.Int64().
					SetValidation(documents.NewElementValidation("length is invalid")),
			),
		documents.Array(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "Post").
			SetTitle("post").
			SetDescription("post").
			AddProperty(
				"id",
				documents.String().
					SetTitle("id").
					SetDescription("id"),
			).
			AddProperty(
				"user",
				documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "User").
					SetTitle("User").
					SetDescription("User model").
					AddProperty(
						"id",
						documents.String().
							SetTitle("Id").
							SetDescription("Id"),
					).
					AddProperty(
						"createAt",
						documents.DateTime().
							SetTitle("create time").
							SetDescription("create time"),
					).
					AddProperty(
						"nickname",
						documents.String().
							SetTitle("nickname").
							SetDescription("nickname"),
					).
					AddProperty(
						"mobile",
						documents.String().
							SetTitle("mobile").
							SetDescription("mobile"),
					).
					AddProperty(
						"gender",
						documents.String().
							SetTitle("gender").
							SetDescription("gender").
							AddEnum("F(female)", "M(male)", "N(unknown)"),
					).
					AddProperty(
						"birthday",
						documents.Date().
							SetTitle("birthday").
							SetDescription("birthday"),
					).
					AddProperty(
						"avatar",
						documents.Struct("github.com/aacfactory/fns-example/standalone/repositories/postgres", "Avatar").
							SetTitle("Avatar").
							SetDescription("Avatar info").
							AddProperty(
								"schema",
								documents.String().
									SetTitle("http schema").
									SetDescription("http schema"),
							).
							AddProperty(
								"domain",
								documents.String().
									SetTitle("domain").
									SetDescription("domain"),
							).
							AddProperty(
								"path",
								documents.String().
									SetTitle("uri path").
									SetDescription("uri path"),
							).
							AddProperty(
								"mimeType",
								documents.String().
									SetTitle("mime type").
									SetDescription("mime type"),
							).
							AddProperty(
								"url",
								documents.String().
									SetTitle("url").
									SetDescription("full url"),
							).
							SetTitle("user avatar").
							SetDescription("user avatar"),
					).
					AddProperty(
						"parent",
						documents.Ref("github.com/aacfactory/fns-example/standalone/modules/users", "User").
							SetTitle("user parent").
							SetDescription("user parent"),
					).
					AddProperty(
						"bd",
						documents.Date(),
					).
					SetTitle("author").
					SetDescription("author"),
			).
			AddProperty(
				"createAt",
				documents.DateTime().
					SetTitle("create time").
					SetDescription("create time"),
			).
			AddProperty(
				"title",
				documents.String().
					SetTitle("title").
					SetDescription("title"),
			).
			AddProperty(
				"content",
				documents.String().
					SetTitle("content").
					SetDescription("content"),
			).
			AddProperty(
				"comments",
				documents.Array(documents.Struct("github.com/aacfactory/fns-example/standalone/modules/posts", "Comment").
					SetTitle("comment").
					SetDescription("comment").
					AddProperty(
						"id",
						documents.Int64().
							SetTitle("id").
							SetDescription("id"),
					).
					AddProperty(
						"postId",
						documents.String().
							SetTitle("post id").
							SetDescription("post id"),
					).
					AddProperty(
						"user",
						documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "User").
							SetTitle("User").
							SetDescription("User model").
							AddProperty(
								"id",
								documents.String().
									SetTitle("Id").
									SetDescription("Id"),
							).
							AddProperty(
								"createAt",
								documents.DateTime().
									SetTitle("create time").
									SetDescription("create time"),
							).
							AddProperty(
								"nickname",
								documents.String().
									SetTitle("nickname").
									SetDescription("nickname"),
							).
							AddProperty(
								"mobile",
								documents.String().
									SetTitle("mobile").
									SetDescription("mobile"),
							).
							AddProperty(
								"gender",
								documents.String().
									SetTitle("gender").
									SetDescription("gender").
									AddEnum("F(female)", "M(male)", "N(unknown)"),
							).
							AddProperty(
								"birthday",
								documents.Date().
									SetTitle("birthday").
									SetDescription("birthday"),
							).
							AddProperty(
								"avatar",
								documents.Struct("github.com/aacfactory/fns-example/standalone/repositories/postgres", "Avatar").
									SetTitle("Avatar").
									SetDescription("Avatar info").
									AddProperty(
										"schema",
										documents.String().
											SetTitle("http schema").
											SetDescription("http schema"),
									).
									AddProperty(
										"domain",
										documents.String().
											SetTitle("domain").
											SetDescription("domain"),
									).
									AddProperty(
										"path",
										documents.String().
											SetTitle("uri path").
											SetDescription("uri path"),
									).
									AddProperty(
										"mimeType",
										documents.String().
											SetTitle("mime type").
											SetDescription("mime type"),
									).
									AddProperty(
										"url",
										documents.String().
											SetTitle("url").
											SetDescription("full url"),
									).
									SetTitle("user avatar").
									SetDescription("user avatar"),
							).
							AddProperty(
								"parent",
								documents.Ref("github.com/aacfactory/fns-example/standalone/modules/users", "User").
									SetTitle("user parent").
									SetDescription("user parent"),
							).
							AddProperty(
								"bd",
								documents.Date(),
							).
							SetTitle("author").
							SetDescription("author"),
					).
					AddProperty(
						"createAt",
						documents.DateTime().
							SetTitle("create time").
							SetDescription("create time"),
					).
					AddProperty(
						"content",
						documents.String().
							SetTitle("content").
							SetDescription("content"),
					)).
					SetTitle("comments").
					SetDescription("latest 10 comments"),
			).
			AddProperty(
				"likes",
				documents.Int64().
					SetTitle("likes").
					SetDescription("likes num"),
			)).
			SetPath("github.com/aacfactory/fns-example/standalone/modules/posts").
			SetName("Posts").
			SetTitle("post list").
			SetDescription("post list"),
		[]documents.FnError{
			{
				Name_: "posts_list_failed",
				Descriptions_: map[string]string{
					"en": "posts list failed",
				},
			},
		},
	)

	doc = document
	return

}
