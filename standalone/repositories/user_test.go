package repositories_test

import (
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/postgres"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-contrib/databases/sql/dac/groups"
	"github.com/aacfactory/fns-example/standalone/repositories"
	"github.com/aacfactory/fns/commons/times"
	"github.com/aacfactory/fns/services/authorizations"
	"github.com/aacfactory/fns/tests"
	"testing"
	"time"
)

func TestUser_Query(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows, queryErr := postgres.Query[repositories.UserRow](
		tests.TODO(),
		0, 10,
		postgres.Conditions(postgres.Eq("Gender", "F")),
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

func TestUser_Insert(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	row, ok, insertErr := postgres.Insert[repositories.UserRow](
		ctx,
		repositories.UserRow{
			Nickname: "u1",
			Mobile:   "14400000001",
			Gender:   "F",
			Birthday: time.Now(),
			Avatar: sql.NullJson[repositories.Avatar]{
				Valid: true,
				E: repositories.Avatar{
					Schema:   "s",
					Domain:   "d",
					Path:     "p",
					MimeType: "m",
					URL:      "u",
				},
			},
			BD: times.DateNow(),
			BT: times.TimeNow(),
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

func TestUser_InsertOrUpdate(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	ctx := authorizations.With(tests.TODO(), authorizations.Authorization{
		Id:         authorizations.StringId([]byte("abc")),
		Account:    nil,
		Attributes: nil,
		ExpireAT:   time.Now().AddDate(1, 1, 1),
	})
	user := repositories.UserRow{
		Nickname: "u1",
		Mobile:   "14400000001",
		Gender:   "F",
		Birthday: time.Now(),
		Avatar:   sql.NullJson[repositories.Avatar]{},
		BD:       times.DateNow(),
		BT:       times.TimeNow(),
	}
	user.Id = "1"
	row, ok, insertErr := postgres.InsertOrUpdate[repositories.UserRow](
		ctx,
		user,
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

func TestUserGenderCount_Query(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	rows, queryErr := postgres.Views[repositories.UserGenderCount](
		tests.TODO(),
		0, 10,
		postgres.GroupBy(groups.Group("Gender")),
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

func TestUserRow_Count(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	count, countErr := postgres.Count[repositories.UserRow](
		tests.TODO(),
		postgres.Eq("Version", 1),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if countErr != nil {
		t.Errorf("%+v", countErr)
		return
	}
	fmt.Println(count)
}

func TestUserRow_Exist(t *testing.T) {
	setupPostgres(t)
	defer tests.Teardown()
	beg := time.Now()
	exist, existErr := postgres.Exist[repositories.UserRow](
		tests.TODO(),
		postgres.Eq("Version", 1),
	)
	fmt.Println("latency", time.Now().Sub(beg).String())
	if existErr != nil {
		t.Errorf("%+v", existErr)
		return
	}
	fmt.Println(exist)
}
