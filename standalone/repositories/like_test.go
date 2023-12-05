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
	row, ok, execErr := postgres.Insert[repositories.PostLikeRow](
		tests.TODO(),
		repositories.PostLikeRow{
			Id:     0,
			PostId: "x",
			UserId: "x",
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_InsertMulti(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows := []repositories.PostLikeRow{
		{0, "p1", "u1"},
		{0, "p2", "u2"},
		{0, "p3", "u3"},
	}
	affected, execErr := postgres.InsertMulti[repositories.PostLikeRow](
		tests.TODO(),
		rows,
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}
	fmt.Println(affected, affected == int64(len(rows)))
	for _, row := range rows {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_Update(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	row, ok, execErr := postgres.Update[repositories.PostLikeRow](
		tests.TODO(),
		repositories.PostLikeRow{
			Id:     1,
			PostId: "x",
			UserId: "x",
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_Delete(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	row, ok, execErr := postgres.Delete[repositories.PostLikeRow](
		tests.TODO(),
		repositories.PostLikeRow{
			Id:     21,
			PostId: "x",
			UserId: "x",
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPostLikeRow_DeleteByCond(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	affected, execErr := postgres.DeleteByCondition[repositories.PostLikeRow](
		tests.TODO(),
		postgres.Eq("PostId", "x"),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}
	fmt.Println(affected)
}
