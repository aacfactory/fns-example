package users_test

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"testing"
)

func Test_P(t *testing.T)  {

	p := make([]byte, 3)
	p[0] = '-'
	copy(p[1:], []byte("{}"))

	fmt.Println(string(p), len(p), string(p[0]), string(p[1]))

}

func Test_Case(t *testing.T)  {

	c := cases.Title(language.English)

	fmt.Println(c.String("Test_Case"))
	fmt.Println(c.String("create"))
	fmt.Println(c.String("createBar"))
	fmt.Println(c.Span([]byte("CreateBar"), true))
	fmt.Println(c.String("createbar"))

}

type Foo struct {
	Key string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	Sub *Foo `json:"sub,omitempty"`
}


func Test_JsonMerge(t *testing.T)  {


	foo := Foo{
		Key:  "a",
		Name: "a",
		Sub:  &Foo{
			Key:  "b",
			Name: "b",
		},
	}

	bar := Foo{
		Key:  "a",
		Sub:  &Foo{
			Key: "c",
			Sub: &Foo{
				Key:  "c",
				Name: "c",
			},
		},
	}

	p1, _ := json.Marshal(foo)
	fmt.Println("p1", string(p1))

	p2, _ := json.Marshal(bar)
	fmt.Println("p1", string(p2))

	r2 := gjson.ParseBytes(p2)


	r2.ForEach(func(key, value gjson.Result) (ok bool) {
		p1 = merge(p1, key.String(), value)
		ok = true
		return
	})

	fmt.Println("=>", string(p1))
}

func merge(dst []byte, srcKey string, srcValue gjson.Result) (result []byte) {
	switch srcValue.Type {
	case gjson.String, gjson.Number, gjson.True, gjson.False:
		affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
		if setErr != nil {
			result = dst
			return
		}
		result = affected
	case gjson.JSON:
		if srcValue.IsArray() {
			affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
			if setErr != nil {
				result = dst
				return
			}
			result = affected
			return
		}
		if srcValue.IsObject() {
			dstSub := gjson.GetBytes(dst, srcKey)
			if !dstSub.Exists() {
				affected, setErr := sjson.SetRawBytes(dst, srcKey, []byte(srcValue.Raw))
				if setErr != nil {
					result = dst
					return
				}
				result = affected
				return
			}

			dstSubRas := []byte(dstSub.Raw)
			srcValue.ForEach(func(key, value gjson.Result) bool {
				dstSubRas = merge(dstSubRas, key.Str, value)
				return true
			})

			affected, setErr := sjson.SetRawBytes(dst, srcKey, dstSubRas)
			if setErr != nil {
				result = dst
				return
			}
			result = affected

		}
	default:
		result = dst
	}

	return
}