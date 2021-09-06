package users

import (
	"fmt"
	"github.com/aacfactory/errors"
	"github.com/aacfactory/fns"
)

func tokenRevoke(ctx fns.Context) (err errors.CodeError) {

	ctx.App().Log().Debug().Message(fmt.Sprintf("%s", ctx.User().String()))

	err = ctx.App().Authorizations().Revoke(ctx)

	return
}
