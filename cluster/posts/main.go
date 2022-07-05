package main

import (
	"fmt"

	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/cluster/posts/modules"
)

var (
	Version string = "v0.0.1"
)

//go:generate fnc codes
func main() {
	// set system environment to make config be active, e.g.: export FNS-ACTIVE=local
	app := fns.New(
		fns.Version(Version),
	)

	if deployErr := app.Deploy(modules.Services()...); deployErr != nil {
		if deployErr != nil {
			app.Log().Error().Caller().Message(fmt.Sprintf("%+v", deployErr))
			return
		}
	}

	if runErr := app.Run(); runErr != nil {
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", runErr))
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Caller().Message("running...")
	}

	if syncErr := app.Sync(); syncErr != nil {
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", syncErr))
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Message("stopped!!!")
	}

}
