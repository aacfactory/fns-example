package main

import (
	"github.com/aacfactory/fns"
	_ "github.com/aacfactory/fns-contrib/databases/redis/clusters"
	"github.com/aacfactory/fns-example/cluster/posts/modules"
	"github.com/aacfactory/fns/context"
	"github.com/aacfactory/fns/transports/middlewares/cors"
)

var (
	// Version
	// go build -ldflags "-X main.Version=${VERSION}" -o standalone
	Version string = "v0.0.1"
)

//go:generate go run -mod=mod github.com/aacfactory/fns-example/cluster/posts/internal/generator -v .
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
