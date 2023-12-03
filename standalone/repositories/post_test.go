package repositories_test

import (
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	_ "github.com/aacfactory/fns-contrib/databases/postgres/dialect"
	"github.com/aacfactory/fns-contrib/databases/sql/dac/specifications"
	"github.com/aacfactory/fns-example/standalone/repositories"
	"github.com/aacfactory/fns/tests"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func setupPostgres(t *testing.T) {
	err := tests.Setup(postgres.New(), tests.WithConfigActive("private"))
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestPostLikeRow_Query(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	ctx := tests.TODO()
	row := repositories.UserRow{}
	spec, specErr := specifications.GetSpecification(ctx, row)
	if specErr != nil {
		t.Errorf("%+v", specErr)
		return
	}
	fmt.Println(spec.String())

	beg := time.Now()

	rows, queryErr := postgres.Query[repositories.UserRow](
		tests.TODO(),
		0, 10,
		postgres.Conditions(postgres.Eq("Gender", "M")),
	)

	fmt.Println("latency", time.Now().Sub(beg).String())

	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	for _, userRow := range rows {
		fmt.Println(fmt.Sprintf("%+v", userRow))
	}
}
