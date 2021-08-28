package main

import (
	"context"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/modules/users"
)

func main() {

	app, appErr := fns.New(
		fns.ConfigRetriever("./config", "YAML", "debug", "app", '-'),
		fns.SecretKeyFile("./config/sk.txt"),
		fns.Version("v0.0.1"),
	)

	if appErr != nil {
		panic(appErr)
		return
	}

	_ = app.Deploy(
		&users.Service{},
	)

	runErr := app.Run(context.TODO())

	if runErr != nil {
		app.Log().Error().Cause(runErr).Caller().Message("app run failed")
		return
	}

	if app.Log().DebugEnabled() {
		app.Log().Debug().Caller().Message("running...")
	}

	app.Sync()

	if app.Log().DebugEnabled() {
		app.Log().Debug().Message("stopped!!!")
	}

}
