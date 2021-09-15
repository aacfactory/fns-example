package repository

import "time"

type UserRow struct {
	Id         string       `col:"ID,PK"`
	CreateBY   string       `col:"CREATE_BY,CREATE_BY"`
	CreateAT   time.Time    `col:"CREATE_AT,CREATE_AT"`
	ModifyBY   string       `col:"MODIFY_BY,MODIFY_BY"`
	ModifyAT   time.Time    `col:"MODIFY_AT,MODIFY_AT"`
	DeleteBY   string       `col:"DELETE_BY,DELETE_BY"`
	DeleteAT   time.Time    `col:"DELETE_AT,DELETE_AT"`
	Version    int64        `col:"VERSION,VERSION"`
	Name       string       `col:"NAME"`
	Password   string       `col:"PASSWORD"`
	Gender     string       `col:"GENDER"`
	Age        int          `col:"AGE"`
	Active     bool         `col:"ACTIVE"`
	SignUpTime time.Time    `col:"SIGN_UP_TIME"`
	Profile    *UserProfile `col:"PROFILE,JSON"`
	Score      float64      `col:"SCORE"`
	DOB        time.Time    `col:"DOB"`
	Posts      []*PostRow   `col:"-,LK" sort:"CREATE_AT DESC"`
}

func (r *UserRow) Table() (string, string, string) {
	return "FNS", "USER", "U"
}

type UserProfile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
