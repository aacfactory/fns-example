package main

import (
	"context"
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
	fns.
		New(
			fns.Version(Version),
		).
		Deploy(modules.Services()...).
		Run(context.Background()).
		Sync()
	return
}
