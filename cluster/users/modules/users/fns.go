// NOTE: This file has been automatically generated by github.com/aacfactory/gcg. Dont Edit it.

package users

import (
	"context"

	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns/service"
	"github.com/aacfactory/fns/service/documents"
	"github.com/aacfactory/fns/service/validators"
)

const (
	_name     = "users"
	_createFn = "create"
	_getFn    = "get"
)

func Create(ctx context.Context, argument CreateArgument) (result *CreateResult, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.NotFound("endpoint was not found").WithMeta("name", _name)
		return
	}
	fr := endpoint.Request(ctx, _createFn, service.NewArgument(argument))
	handled := CreateResult{}
	hasResult, handleErr := fr.Get(ctx, &handled)
	if handleErr != nil {
		err = handleErr
		return
	}
	if hasResult {
		result = &handled
	}
	return
}

func Get(ctx context.Context, argument GetArgument) (result *User, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.NotFound("endpoint was not found").WithMeta("name", _name)
		return
	}
	fr := endpoint.Request(ctx, _getFn, service.NewArgument(argument))
	handled := User{}
	hasResult, handleErr := fr.Get(ctx, &handled)
	if handleErr != nil {
		err = handleErr
		return
	}
	if hasResult {
		result = &handled
	}
	return
}

func Service() (svc service.Service) {
	components := []service.Component{}
	svc = &_service_{
		Abstract: service.NewAbstract(
			_name,
			false,
			components...,
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
		// argument
		arg := CreateArgument{}
		scanErr := argument.As(&arg)
		if scanErr != nil {
			err = errors.BadRequest("users: scan request argument failed").WithCause(scanErr).WithMeta("service", _name).WithMeta("fn", _createFn)
			return
		}
		validateErr := validators.Validate(arg)
		if validateErr != nil {
			err = errors.BadRequest("users: invalid request argument").WithMeta("service", _name).WithMeta("fn", _createFn).WithCause(validateErr)
			return
		}
		// sql begin transaction
		beginTransactionErr := sql.BeginTransaction(ctx)
		if beginTransactionErr != nil {
			err = errors.ServiceError("users: begin sql transaction failed").WithMeta("service", _name).WithMeta("fn", _createFn).WithCause(beginTransactionErr)
			return
		}
		// handle
		v, err = create(ctx, arg)
		// sql close transaction
		if err == nil {
			commitTransactionErr := sql.CommitTransaction(ctx)
			if commitTransactionErr != nil {
				err = errors.ServiceError("users: commit sql transaction failed").WithMeta("service", _name).WithMeta("fn", _createFn).WithCause(commitTransactionErr)
				_ = sql.RollbackTransaction(ctx)
				return
			}
		}
		break
	case _getFn:
		// argument
		arg := GetArgument{}
		scanErr := argument.As(&arg)
		if scanErr != nil {
			err = errors.BadRequest("users: scan request argument failed").WithCause(scanErr).WithMeta("service", _name).WithMeta("fn", _getFn)
			return
		}
		validateErr := validators.Validate(arg)
		if validateErr != nil {
			err = errors.BadRequest("users: invalid request argument").WithMeta("service", _name).WithMeta("fn", _getFn).WithCause(validateErr)
			return
		}
		// handle
		v, err = get(ctx, arg)
		break
	default:
		err = errors.NotFound("users: fn was not found").WithMeta("service", _name).WithMeta("fn", fn)
		break
	}
	return

}

func (svc *_service_) Document() (doc service.Document) {
	sd := documents.NewService(_name, "User service")
	sd.AddFn(
		"create", "Create user", "Create a user\n----------\nerrors:\n| Name                     | Code    | Description                   |\n|--------------------------|---------|-------------------------------|\n| users_create_failed      | 500     | create user failed            |", false, false,
		documents.Struct("main/modules/users", "CreateArgument", "Create user argument", "Create user argument").
			AddProperty("nickname",
				documents.String().SetTitle("Nickname").SetDescription("Nickname"),
			).
			AddProperty("mobile",
				documents.String().SetTitle("Mobile").SetDescription("Mobile"),
			).
			AddProperty("gender",
				documents.String().SetTitle("Gender").SetDescription("Gender").AddEnum("F(female)", "M(male)", "N(unknown)"),
			).
			AddProperty("birthday",
				documents.Date().SetTitle("Birthday").SetDescription("Birthday").AsRequired(`validate:"required" message:"birthday is invalid"`),
			),
		documents.Struct("main/modules/users", "CreateResult", "Create user result", "Create user result").
			AddProperty("id",
				documents.String().SetTitle("id").SetDescription("user id"),
			).
			AddProperty("token",
				documents.String().SetTitle("token").SetDescription("user token"),
			),
	)
	sd.AddFn(
		"get", "Get", "Get a user\n----------\nerrors:\n| Name                     | Code    | Description                   |\n|--------------------------|---------|-------------------------------|\n| users_get_failed         | 500     | get user failed            |\n| users_get_nothing        | 404     | user was not found            |", false, false,
		documents.Struct("main/modules/users", "GetArgument", "Get user argument", "Get user argument").
			AddProperty("id",
				documents.String().SetTitle("Id").SetDescription("Id").AsRequired(`validate:"required" message:"id is invalid"`),
			),
		documents.Struct("main/modules/users", "User", "User", "User model").
			AddProperty("id",
				documents.String().SetTitle("Id").SetDescription("Id"),
			).
			AddProperty("createAt",
				documents.DateTime().SetTitle("create time").SetDescription("create time"),
			).
			AddProperty("nickname",
				documents.String().SetTitle("nickname").SetDescription("nickname"),
			).
			AddProperty("mobile",
				documents.String().SetTitle("mobile").SetDescription("mobile"),
			).
			AddProperty("gender",
				documents.String().SetTitle("gender").SetDescription("gender").AddEnum("F(female)", "M(male)", "N(unknown)"),
			).
			AddProperty("birthday",
				documents.Date().SetTitle("birthday").SetDescription("birthday"),
			).
			AddProperty("avatar",
				documents.Struct("main/repository", "Avatar", "Avatar", "Avatar info").
					AddProperty("schema",
						documents.String().SetTitle("http schema").SetDescription("http schema"),
					).
					AddProperty("domain",
						documents.String().SetTitle("domain").SetDescription("domain"),
					).
					AddProperty("path",
						documents.String().SetTitle("uri path").SetDescription("uri path"),
					).
					AddProperty("mimeType",
						documents.String().SetTitle("mime type").SetDescription("mime type"),
					).
					AddProperty("url",
						documents.String().SetTitle("url").SetDescription("full url"),
					).SetTitle("user avatar").SetDescription("user avatar"),
			),
	)
	doc = sd
	return
}
