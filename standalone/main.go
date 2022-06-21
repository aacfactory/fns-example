package main

import (
	"fmt"
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/standalone/modules/hellos"
	"github.com/aacfactory/fns-example/standalone/modules/users"
	"github.com/aacfactory/fns/service/builtin/authorizations"
	_ "github.com/lib/pq"
)

//go:generate fnc codes .

//main
func main() {

	app := fns.New(
		//fns.ConfigRetriever("./config", "YAML", fns.ConfigActiveFromENV("FNS-ACTIVE"), "app", '-'),
		//fns.Server(http3.Server()),
		fns.ExtraListeners(),
		fns.Version("v0.0.1"),
	)

	deployErr := app.Deploy(
		authorizations.Service(),
		hellos.Service(),
		users.Service(),
	)

	if deployErr != nil {
		//app.Log().Error().Cause(deployErr).Caller().Message("app deploy failed")
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", deployErr))
		return
	}
	if runErr := app.Run(); runErr != nil {
		app.Log().Error().Caller().Message(fmt.Sprintf("%+v", runErr))
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
