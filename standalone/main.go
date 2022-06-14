package main

import (
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/standalone/modules/hellos"
	_ "github.com/lib/pq"
)

func main() {

	app := fns.New(
		fns.ConfigRetriever("./config", "YAML", fns.ConfigActiveFromENV("FNS-ACTIVE"), "app", '-'),
		fns.SecretKey("test"),
		fns.MAXPROCS(0, 16),
		fns.Version("v0.0.1"),
	)

	deployErr := app.Deploy(
		hellos.Service(),
	)

	if deployErr != nil {
		app.Log().Error().Cause(deployErr).Caller().Message("app deploy failed")
		return
	}
	runErr := app.Run()
	if runErr != nil {
		app.Log().Error().Cause(runErr).Caller().Message("app run failed")
		return
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Caller().Message("running...")
	}
	syncErr := app.Sync()
	if syncErr != nil {
		app.Log().Error().Cause(syncErr).Caller().Message("app sync failed")
		return
	}
	if app.Log().DebugEnabled() {
		app.Log().Debug().Message("stopped!!!")
	}
}
