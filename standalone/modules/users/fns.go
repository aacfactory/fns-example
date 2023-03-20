// NOTE: this file has been automatically generated, DON'T EDIT IT!!!

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

func Create(ctx context.Context, argument CreateArgument) (result CreateArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("users: endpoint was not found").WithMeta("name", _name)
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
		err = errors.Warning("users: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _createFn).
			WithCause(scanErr)
		return
	}
	return
}

func Get(ctx context.Context, argument GetArgument) (result GetArgument, err errors.CodeError) {
	endpoint, hasEndpoint := service.GetEndpoint(ctx, _name)
	if !hasEndpoint {
		err = errors.Warning("users: endpoint was not found").WithMeta("name", _name)
		return
	}
	fr, requestErr := endpoint.RequestSync(ctx, service.NewRequest(ctx, _name, _getFn, service.NewArgument(argument)))
	if requestErr != nil {
		err = requestErr
		return
	}
	if !fr.Exist() {
		return
	}
	scanErr := fr.Scan(&result)
	if scanErr != nil {
		err = errors.Warning("users: scan future result failed").
			WithMeta("service", _name).WithMeta("fn", _getFn).
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
			err = errors.Warning("users: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// sql begin transaction
		beginTransactionErr := sql.BeginTransaction(ctx)
		if beginTransactionErr != nil {
			err = errors.Warning("users: begin sql transaction failed").WithCause(beginTransactionErr)
			return
		}
		// execute function
		v, err = create(ctx, param)
		// sql commit transaction
		if err == nil {
			commitTransactionErr := sql.CommitTransaction(ctx)
			if commitTransactionErr != nil {
				_ = sql.RollbackTransaction(ctx)
				err = errors.ServiceError("users: commit sql transaction failed").WithCause(commitTransactionErr)
				return
			}
		}
		break
	case _getFn:
		// param
		param := GetArgument{}
		paramErr := argument.As(&param)
		if paramErr != nil {
			err = errors.Warning("users: decode request argument failed").WithCause(paramErr)
			break
		}
		err = validators.ValidateWithErrorTitle(param, "invalid")
		if err != nil {
			break
		}
		// execute function
		v, err = get(ctx, param)
		break
	default:
		err = errors.Warning("users: fn was not found").WithMeta("service", _name).WithMeta("fn", fn)
		break
	}
	return
}

func (svc *_service_) Document() (doc service.Document) {
	document := documents.NewService(_name, "User service")
	// create
	document.AddFn(
		"create", "Create user", "Create a user", false, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "CreateArgument").
			SetTitle("Create user argument").
			SetDescription("Create user argument").
			AddProperty(
				"nickname",
				documents.String().
					SetTitle("Nickname").
					SetDescription("Nickname").
					SetValidation(documents.NewElementValidation("nickname is invalid")),
			).
			AddProperty(
				"mobile",
				documents.String().
					SetTitle("Mobile").
					SetDescription("Mobile").
					SetValidation(documents.NewElementValidation("mobile is invalid")),
			).
			AddProperty(
				"gender",
				documents.String().
					SetTitle("Gender").
					SetDescription("Gender").
					AddEnum("F(female)", "M(male)", "N(unknown)").
					SetValidation(documents.NewElementValidation("gender is invalid")),
			).
			AddProperty(
				"birthday",
				documents.Date().
					SetTitle("Birthday").
					SetDescription("Birthday").
					AsRequired().
					SetValidation(documents.NewElementValidation("birthday is invalid")),
			),
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "CreateResult").
			SetTitle("Create user result").
			SetDescription("Create user result").
			AddProperty(
				"id",
				documents.String().
					SetTitle("id").
					SetDescription("user id"),
			).
			AddProperty(
				"token",
				documents.String().
					SetTitle("token").
					SetDescription("user token"),
			),
		[]documents.FnError{
			{
				Name_: "users_create_failed",
				Descriptions_: map[string]string{
					"en": "users create failed",
				},
			},
		},
	)

	// get
	document.AddFn(
		"get", "Get", "Get a user", false, false,
		documents.Struct("github.com/aacfactory/fns-example/standalone/modules/users", "GetArgument").
			SetTitle("Get user argument").
			SetDescription("Get user argument").
			AddProperty(
				"id",
				documents.String().
					SetTitle("Id").
					SetDescription("Id").
					AsRequired().
					SetValidation(documents.NewElementValidation("id is invalid")),
			),
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
			),
		[]documents.FnError{
			{
				Name_: "users_get_failed",
				Descriptions_: map[string]string{
					"en": "users get failed",
				},
			},
			{
				Name_: "users_get_nothing",
				Descriptions_: map[string]string{
					"en": "users get nothing",
				},
			},
		},
	)

	doc = document
	return

}
