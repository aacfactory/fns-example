package main_test

import (
	"context"
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns/cmd/generates"
	"path/filepath"
	"testing"
)

func TestF(t *testing.T) {
	fmt.Println(filepath.Abs("../../"))
	g := generates.New(
		generates.WithAnnotations(postgres.FAG()...),
	)
	if err := g.Execute(context.Background(), []string{"-v", "../../"}...); err != nil {
		fmt.Println(fmt.Sprintf("%+v", err))
	}
}
