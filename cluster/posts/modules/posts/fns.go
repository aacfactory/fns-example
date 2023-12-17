// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

package posts

import (
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/runtime"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/services/commons"
	"github.com/aacfactory/fns/services/documents"
	"github.com/aacfactory/fns/services/validators"
	"github.com/aacfactory/fns/transports/middlewares/cachecontrol"
)

var (
	_endpointName = []byte("posts")
	_listFnName   = []byte("list")
)

// +-------------------------------------------------------------------------------------------------------------------+

func List(ctx context.Context, param ListParam) (result Posts, err error) {
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// handle
	eps := runtime.Endpoints(ctx)
	response, handleErr := eps.Request(ctx, _endpointName, _listFnName, param)
	if handleErr != nil {
		err = handleErr
		return
	}
	result, err = services.ValueOfResponse[Posts](response)
	return

}

func _list(ctx services.Request) (v any, err error) {
	// param
	param, paramErr := services.ValueOfParam[ListParam](ctx.Param())
	if paramErr != nil {
		err = errors.BadRequest("scan params failed").WithCause(paramErr)
		return
	}
	// validate param
	if err = validators.ValidateWithErrorTitle(param, "invalid"); err != nil {
		return
	}
	// handle
	v, err = list(ctx, param)
	// cache control
	cachecontrol.Make(ctx, cachecontrol.MaxAge(10), cachecontrol.Public())
	return

}

// +-------------------------------------------------------------------------------------------------------------------+

func Component[C services.Component](ctx context.Context, name string) (component C, has bool) {
	component, has = services.LoadComponent[C](ctx, _endpointName, name)
	return
}

// +-------------------------------------------------------------------------------------------------------------------+

func Service() (v services.Service) {
	v = &_service{
		Abstract: services.NewAbstract(
			string(_endpointName),
			false,
		),
	}
	return
}

// +-------------------------------------------------------------------------------------------------------------------+

type _service struct {
	services.Abstract
}

func (svc *_service) Construct(options services.Options) (err error) {
	if err = svc.Abstract.Construct(options); err != nil {
		return
	}
	svc.AddFunction(commons.NewFn(string(_listFnName), true, false, false, false, true, true, _list))
	return
}

func (svc *_service) Document() (document documents.Endpoint) {
	document = documents.New(svc.Name(), "Posts", "posts")
	// list
	document.AddFn(
		documents.NewFn("list").
			SetInfo("list", "list posts").
			SetReadonly(true).SetInternal(false).SetDeprecated(false).
			SetAuthorization(false).SetPermission(false).
			SetParam(documents.Struct("github.com/aacfactory/fns-example/cluster/posts/modules/posts", "ListParam").
				AddProperty(
					"userId",
					documents.String().
						AsRequired().
						SetValidation(documents.NewValidation("user_id_required")),
				).
				AddProperty(
					"offset",
					documents.Int64(),
				).
				AddProperty(
					"length",
					documents.Int64(),
				)).
			SetResult(documents.Array(documents.Struct("github.com/aacfactory/fns-example/cluster/posts/modules/posts", "Post").
				AddProperty(
					"id",
					documents.String(),
				).
				AddProperty(
					"userId",
					documents.String(),
				).
				AddProperty(
					"createAT",
					documents.DateTime(),
				).
				AddProperty(
					"title",
					documents.String(),
				).
				AddProperty(
					"content",
					documents.String(),
				).
				AddProperty(
					"likes",
					documents.Int64(),
				).
				AddProperty(
					"comments",
					documents.Array(documents.Struct("github.com/aacfactory/fns-example/cluster/posts/modules/posts", "Comment").
						AddProperty(
							"id",
							documents.Int64(),
						).
						AddProperty(
							"postId",
							documents.String(),
						).
						AddProperty(
							"userId",
							documents.String(),
						).
						AddProperty(
							"createAT",
							documents.DateTime(),
						).
						AddProperty(
							"content",
							documents.String(),
						)),
				)).
				SetPath("github.com/aacfactory/fns-example/cluster/posts/modules/posts").
				SetName("Posts")).
			SetErrors("posts_not_found\nzh: zh_message\nen: en_message"),
	)
	return
}
