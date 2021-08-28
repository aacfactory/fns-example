package users_test

import (
	"fmt"
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