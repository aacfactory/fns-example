package users_test

import (
	"fmt"
	"github.com/aacfactory/json"
	"testing"
	"time"
)

func TestDur(t *testing.T) {
	d := 1 * time.Second
	p, err := json.Marshal(d)
	fmt.Println(string(p), err)
}
