package main

import (
	"context"
	"fmt"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/standalone/modules"
)

var (
	// Version
	// go build -ldflags "-X main.Version=${VERSION}" -o standalone
	Version string = "v0.0.1"
)

//go:generate fnc codes --debug .
func main() {
	// set system environment to make config be active, e.g.: export FNS-ACTIVE=local
	app := fns.New(
		fns.Version(Version),
	)
	// deploy services
	if deployErr := app.Deploy(modules.Services()...); deployErr != nil {
		if deployErr != nil {
			app.Log().Error().Caller().Message(fmt.Sprintf("%+v", deployErr))
			return
		}
	}
	// run
	if runErr := app.Run(context.TODO()); runErr != nil {
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", runErr))
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Caller().Message("running...")
	}
	// sync signals
	if syncErr := app.Sync(); syncErr != nil {
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", syncErr))
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Message("stopped!!!")
	}
	return
}
