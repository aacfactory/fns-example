package repositories_test

import (
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	_ "github.com/aacfactory/fns-contrib/databases/postgres/dialect"
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
	beg := time.Now()
	rows, queryErr := postgres.Query[repositories.PostLikeRow](
		tests.TODO(),
		0, 10,
		postgres.Conditions(postgres.Eq("PostId", "1")),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	for _, row := range rows {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_One(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	row, has, queryErr := postgres.One[repositories.PostLikeRow](
		tests.TODO(),
		postgres.Conditions(postgres.Eq("PostId", "1")),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	if has {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_All(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows, queryErr := postgres.ALL[repositories.PostLikeRow](
		tests.TODO(),
		postgres.Conditions(postgres.Eq("PostId", "1")),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	for _, row := range rows {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_Insert(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows, queryErr := postgres.ALL[repositories.PostLikeRow](
		tests.TODO(),
		postgres.Conditions(postgres.Eq("PostId", "1")),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	for _, row := range rows {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}
