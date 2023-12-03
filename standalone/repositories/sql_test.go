package repositories_test

import (
	stdsql "database/sql"
	"encoding/json"
	"fmt"
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-example/standalone/repositories"
	"github.com/aacfactory/fns/commons/times"
	"github.com/aacfactory/fns/commons/uid"
	"github.com/aacfactory/fns/services"
	"github.com/aacfactory/fns/tests"
	_ "github.com/lib/pq"
	"testing"
	"time"
)

func setup(t *testing.T) {
	err := tests.Setup(sql.New(), tests.WithConfigActive("private"))
	if err != nil {
		t.Fatalf("%+v", err)
	}
}

func TestQuery(t *testing.T) {
	setup(t)
	defer tests.Teardown()
	ctx := tests.TODO()
	query := `SELECT "ID", "CREATE_BY", "CREATE_AT", "MODIFY_BY", "MODIFY_AT", "DELETE_BY", "DELETE_AT", "VERSION", "NICKNAME", "MOBILE", "GENDER", "BIRTHDAY", "AVATAR", "BD", "BT"
	FROM "FNS"."USER" WHERE "GENDER" = $1 AND "VERSION" = $2`
	ctx = services.AcquireRequest(ctx, []byte("foo"), []byte("fn"), nil)
	begErr := sql.Begin(ctx)
	if begErr != nil {
		t.Errorf("%+v", begErr)
		return
	}
	defer sql.Rollback(ctx)
	rows, queryErr := sql.Query(ctx, []byte(query), "F", 1)
	if queryErr != nil {
		t.Errorf("%+v", queryErr)
		return
	}
	defer rows.Close()
	rp, rowsErr := json.Marshal(rows)
	if rowsErr != nil {
		t.Errorf("%+v", rowsErr)
		return
	}
	rows = sql.Rows{}
	rowsErr = json.Unmarshal(rp, &rows)
	if rowsErr != nil {
		t.Errorf("%+v", rowsErr)
		return
	}
	for rows.Next() {
		id := ""
		createBY := ""
		createAT := stdsql.NullTime{}
		modifyBY := ""
		modifyAT := stdsql.NullTime{}
		deleteBY := ""
		deleteAT := stdsql.NullTime{}
		version := int64(0)
		nickname := ""
		mobile := ""
		gender := ""
		birthday := time.Time{}
		avatar := json.RawMessage{}
		db := times.Date{}
		bt := times.Time{}
		scanErr := rows.Scan(
			&id, &createBY, &createAT, &modifyBY, &modifyAT, &deleteBY, &deleteAT, &version,
			&nickname, &mobile, &gender, &birthday, &avatar, &db, &bt,
		)
		if scanErr != nil {
			t.Errorf("%+v", scanErr)
			return
		}
		fmt.Println("-----")
		fmt.Println("id", id)
		fmt.Println("createBY", createBY)
		fmt.Println("createAT", createAT)
		fmt.Println("modifyBY", modifyBY)
		fmt.Println("modifyAT", modifyAT)
		fmt.Println("deleteBY", deleteBY)
		fmt.Println("deleteAT", deleteAT)
		fmt.Println("version", version)
		fmt.Println("nickname", nickname)
		fmt.Println("mobile", mobile)
		fmt.Println("gender", gender)
		fmt.Println("birthday", birthday)
		fmt.Println("avatar", string(avatar))
		fmt.Println("db", db)
		fmt.Println("bt", bt)
	}
}

func TestUpdate(t *testing.T) {
	setup(t)
	defer tests.Teardown()
	ctx := tests.TODO()

	query := `INSERT INTO "FNS"."USER"(
	"ID", "CREATE_BY", "CREATE_AT", "MODIFY_BY", "MODIFY_AT", "DELETE_BY", "DELETE_AT", "VERSION", "NICKNAME", "MOBILE", "GENDER", "BIRTHDAY", "AVATAR", "BD", "BT")
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`

	id := uid.UID()
	arguments := []any{
		id, id, time.Now(), "", stdsql.NullTime{}, "", stdsql.NullTime{}, 1,
		"nickname", "14400000001", "F", time.Now(),
		sql.NullJson[repositories.Avatar]{},
		times.DataOf(time.Now()), times.TimeOf(time.Now()),
	}

	args := sql.Arguments(arguments)
	argsP, encodeErr := json.Marshal(args)
	if encodeErr != nil {
		t.Errorf("%+v", encodeErr)
		return
	}
	fmt.Println(string(argsP))
	args = sql.Arguments{}
	decodeErr := json.Unmarshal(argsP, &args)
	if decodeErr != nil {
		t.Errorf("%+v", decodeErr)
		return
	}
	fmt.Println(fmt.Sprintf("%+v", args))

	ctx = services.AcquireRequest(ctx, []byte("foo"), []byte("fn"), nil)
	begErr := sql.Begin(ctx)
	if begErr != nil {
		t.Errorf("%+v", begErr)
		return
	}
	defer sql.Rollback(ctx)

	result, execErr := sql.Execute(ctx, []byte(query), args...)
	if execErr != nil {
		t.Errorf("%+v", execErr)
		return
	}

	cmtErr := sql.Commit(ctx)
	if cmtErr != nil {
		t.Errorf("%+v", cmtErr)
		return
	}

	fmt.Println(fmt.Sprintf("%+v", result))
}

func TestNull(t *testing.T) {
	p := []byte("null")
	s := ""
	err := json.Unmarshal(p, &s)
	fmt.Println(s, err)
}
