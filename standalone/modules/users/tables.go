package users

import (
	"time"
)

type UserRow struct {
	Id         string       `col:"ID"`
	Name       string       `col:"NAME"`
	Password   string       `col:"PASSWORD"`
	Gender     string       `col:"GENDER"`
	Age        *UserAge     `col:"AGE,FK"`
	Active     bool         `col:"ACTIVE"`
	SignUpTime time.Time    `col:"SIGN_UP_TIME"`
	Profile    []*UserProfile `col:"PROFILE"`
	Score      float64      `col:"SCORE"`
	DOB        time.Time    `col:"DOB"`
}

type UserProfile struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type UserAge struct {
	Id int  `col:"ID,PK"`
	V  string `col:"V"`
}
