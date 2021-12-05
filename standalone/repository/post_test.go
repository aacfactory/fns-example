package repository

import (
	"crypto/md5"
	"fmt"
	"github.com/aacfactory/json"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

type X struct {
	Id string
}

func TestName(t *testing.T) {

	v := make([]*PostRow, 0, 1)

	rt := reflect.ValueOf(v)
	x := rt.Interface()

	fmt.Println(&x)
	fmt.Println(fff(&x))
	fmt.Println(x)

	s := "sss"
	fmt.Println(fmt.Sprintf("%x",md5.Sum([]byte(s))))

	z := X{Id: "sss"}
	p, _ := json.Marshal(z)
	fmt.Println(string(p))

}

func fff(v interface{}) (err error) {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if reflect.TypeOf(rv.Interface()).Kind() != reflect.Slice {
		err = fmt.Errorf("fns SQL Rows: scan failed for target elem is not slice or struct")
		return
	}
	var elemType reflect.Type
	elemIsPtr := false
	rvt := reflect.TypeOf(rv.Interface())
	elem := rvt.Elem()
	fmt.Println(elem)
	if elem.Kind() == reflect.Ptr {
		if elem.Elem().Kind() != reflect.Struct {
			err = fmt.Errorf("fns SQL Rows: scan failed for element of target is not struct or ptr of struct")
			return
		}
		elemIsPtr = true
		elemType = elem.Elem()
	} else if elem.Kind() == reflect.Struct {
		elemIsPtr = false
		elemType = elem
	} else {
		err = fmt.Errorf("fns SQL Rows: scan failed for element of target is not struct or ptr of struct")
		return
	}
	rv0 := reflect.MakeSlice(rvt, 0, 1)
	x := reflect.New(elemType)
	x.Elem().FieldByName("Id").SetString("id")
	if elemIsPtr {
		rv0 = reflect.Append(rv0, x)
	} else {
		rv0 = reflect.Append(rv0, x.Elem())
	}

	fmt.Println(rv0.Interface())
	rv.Set(rv0)
	return
}

func xxx() *UserRow {
	fmt.Println("xxx")
	return &UserRow{
		Id: "fff",
	}
}

func TestXXX(t *testing.T) {
	p := &PostRow{
		Id:       "x",
		CreateBY: "",
		CreateAT: time.Time{},
		ModifyBY: "",
		ModifyAT: time.Time{},
		Version:  0,
		Title:    "",
		Content:  "",
		Author:   nil,
		Likes:    0,
		Comments: nil,
	}
	rv := reflect.Indirect(reflect.ValueOf(p))
	fv := rv.FieldByName("Author")
	ptr := reflect.ValueOf(xxx).Pointer()
	fmt.Println(ptr)
	x := reflect.NewAt(reflect.TypeOf(xxx), unsafe.Pointer(ptr)).Elem()
	fmt.Println(x.CanInterface(), x.CanAddr(), x.CanConvert(reflect.TypeOf(&UserRow{})))

	fv.Set(x.Elem().Call(nil)[0])
	fmt.Println("!")
	fmt.Println(p)
	a := p.Author
	fmt.Println("!")
	a.Id = ""
	fmt.Println("!")
	fmt.Println(a.Id)

}
