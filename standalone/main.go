package main

import (
	"github.com/aacfactory/fns"
	"github.com/aacfactory/fns-example/standalone/modules"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/transports/middlewares/cors"
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
		).
		Deploy(modules.Services()...).
		Run(context.TODO()).
		Sync()
	return
}
