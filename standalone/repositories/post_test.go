package repositories_test

import (
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	_ "github.com/aacfactory/fns-contrib/databases/postgres/dialect"
	"github.com/aacfactory/fns-example/standalone/repositories"
	"github.com/aacfactory/fns/services/authorizations"
	"github.com/aacfactory/fns/tests"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func TestPost_Query(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows, queryErr := postgres.Query[repositories.PostRow](
		tests.TODO(),
		0, 10,
		postgres.Conditions(postgres.Eq("Version", 1)),
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

func TestPost_Insert(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	row, ok, insertErr := postgres.Insert[repositories.PostRow](
		ctx,
		repositories.PostRow{
			CreateAT: time.Time{},
			Version:  0,
			Title:    "test",
			Content:  "test",
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if insertErr != nil {
		t.Errorf("%+v", insertErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}

}

func TestPost_InsertWhenNotExist(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	row, ok, insertErr := postgres.InsertWhenExist[repositories.PostRow](
		ctx,
		repositories.PostRow{
			CreateAT: time.Time{},
			Version:  0,
			Title:    "test",
			Content:  "test",
		},
		postgres.SubQuery(repositories.UserRow{}, "Id", postgres.Eq("Id", "cb1nn3de2f64qo1sfu3g")),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if insertErr != nil {
		t.Errorf("%+v", insertErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}

}

func TestPost_Update(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	user := repositories.UserRow{}
	user.Id = "1"
	row, ok, updateErr := postgres.Update[repositories.PostRow](
		ctx,
		repositories.PostRow{
			Id:       "1",
			User:     user,
			CreateAT: time.Time{},
			Version:  1,
			Title:    "test",
			Content:  "test1",
			Comments: nil,
			Likes:    0,
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if updateErr != nil {
		t.Errorf("%+v", updateErr)
		return
	}
	if ok {
		fmt.Println(fmt.Sprintf("%+v", row))
	}
}

func TestPost_UpdateFields(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	user := repositories.UserRow{}
	user.Id = "3"
	affected, updateErr := postgres.UpdateFields[repositories.PostRow](
		ctx,
		postgres.Field("User", user),
		postgres.Eq("Id", "1"),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if updateErr != nil {
		t.Errorf("%+v", updateErr)
		return
	}
	fmt.Println(affected)
}

func TestPost_Delete(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	row, ok, deleteErr := postgres.Delete[repositories.PostRow](
		ctx,
		repositories.PostRow{
			Id:      "clnhn2qsvgs2te7gu5tg",
			Version: 1,
		},
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if deleteErr != nil {
		t.Errorf("%+v", deleteErr)
		return
	}
	fmt.Println(ok, row)
}

func TestPost_DeleteCond(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	affected, deleteErr := postgres.DeleteByCondition[repositories.PostRow](
		ctx,
		postgres.Like("Title", "post"),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if deleteErr != nil {
		t.Errorf("%+v", deleteErr)
		return
	}
	fmt.Println(affected)
}

func TestPost_QueryIn(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	ctx := tests.TODO()
	beg := time.Now()
	rows, queryErr := postgres.Query[repositories.PostRow](
		ctx,
		0, 10,
		postgres.Conditions(postgres.In("User",
			postgres.SubQuery(repositories.UserRow{}, "Id", postgres.Eq("Id", "cb1nn3de2f64qo1sfu3g")))),
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
