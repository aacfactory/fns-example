package repository

import (
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {

	v := make([]*PostRow, 0, 1)

	rt := reflect.ValueOf(v)
	x := rt.Interface()

	fmt.Println(&x)
	fmt.Println(fff(&x))
	fmt.Println(x)
}

func fff(v interface{}) (err error)  {
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

