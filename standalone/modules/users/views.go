package users

import (
	"github.com/aacfactory/json"
	"time"
)

func UserViewMapFromTableRow(row *UserRow) (u *User) {
	if row == nil {
		return
	}
	u = &User{
		Id:         row.Id,
		Name:       row.Name,
		Password:   row.Password,
		Gender:     row.Gender,
		Age:        row.Age,
		Active:     row.Active,
		SignUpTime: row.SignUpTime,
		Profile:    row.Profile,
		Score:      row.Score,
		DOB:        row.DOB,
	}
	return
}

type User struct {
	Id         string          `json:"id,omitempty"`
	Name       string          `json:"name,omitempty"`
	Password   string          `json:"password,omitempty"`
	Gender     string          `json:"gender,omitempty"`
	Age        int             `json:"age,omitempty"`
	Active     bool            `json:"active,omitempty"`
	SignUpTime time.Time       `json:"signUpTime,omitempty"`
	Profile    json.RawMessage `json:"profile,omitempty"`
	Score      float64         `json:"score,omitempty"`
	DOB        time.Time       `json:"dob,omitempty"`
}
