package main

import (
	"context"
	"fmt"
	"github.com/aacfactory/fns"
	_ "github.com/aacfactory/fns-contrib/authorizations/jwt"
	"github.com/aacfactory/fns-contrib/databases/redis"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-example/standalone/modules/samples"
	"github.com/aacfactory/json"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	x, _ := json.Marshal(time.Now())
	fmt.Println(string(x))

	app, appErr := fns.New(
		fns.ConfigRetriever("./config", "YAML", fns.ConfigActiveFromENV("FNS_ACTIVE"), "app", '-'),
		fns.SecretKeyFile("./config/sk.txt"),
		fns.Version("v0.0.1"),
	)

	if appErr != nil {
		fmt.Printf("%+v\n", appErr)
		panic(appErr)
		return
	}

	_ = app.Deploy(
		sql.Service(),
		redis.Service(),
		samples.Service(),
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
