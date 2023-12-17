package main

import (
	"context"
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns/cmd/generates"
	"os"
)

func main() {
	g := generates.New(
		generates.WithAnnotations(postgres.FAG()...),
	)
	if err := g.Execute(context.Background(), os.Args...); err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
}
