package main

import (
	"github.com/aacfactory/fns"
	_ "github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-contrib/transports/handlers/documents"
	"github.com/aacfactory/fns-contrib/transports/handlers/pprof"
	"github.com/aacfactory/fns-contrib/transports/handlers/websockets"
	"github.com/aacfactory/fns-example/standalone/modules"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/transports/middlewares/compress"
	"github.com/aacfactory/fns/transports/middlewares/cors"
	_ "github.com/lib/pq"
)

var (
	// Version
	// go build -ldflags "-X main.Version=${VERSION}" -o standalone
	Version string = "v0.0.1"
)

//go:generate go run -mod=mod github.com/aacfactory/fns-example/standalone/internal/generator -v .
func main() {
	// set system environment to make config be active, e.g.: export FNS-ACTIVE=local
	fns.
		New(
			fns.Version(Version),
			fns.Middleware(cors.New()),
			fns.Middleware(compress.New()),
			fns.Handler(documents.New()),
			fns.Handler(pprof.New()),
			fns.Handler(websockets.New()),
		).
		Deploy(modules.Services()...).
		Run(context.TODO()).
		Sync()
	return
}
