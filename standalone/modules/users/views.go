package users

import (
	"github.com/aacfactory/json"
	"time"
)

type User struct {
	Id         string          `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Password   string          `json:"password,omitempty"`
	Gender     []byte          `json:"gender,omitempty"`
	Age        int             `json:"age,omitempty"`
	Active     bool            `json:"active,omitempty"`
	SignUpTime time.Time       `json:"signUpTime,omitempty"`
	Profile    json.RawMessage `json:"profile,omitempty"`
	Score      float64         `json:"score,omitempty"`
	DOB        time.Time       `json:"dob,omitempty"`
}
