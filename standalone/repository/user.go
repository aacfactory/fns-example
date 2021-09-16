package repository

import "time"

type UserRow struct {
	Id         string       `col:"ID,PK"`
	CreateBY   string       `col:"CREATE_BY,ACB"`
	CreateAT   time.Time    `col:"CREATE_AT,ACT"`
	ModifyBY   string       `col:"MODIFY_BY,AMB"`
	ModifyAT   time.Time    `col:"MODIFY_AT,AMT"`
	DeleteBY   string       `col:"DELETE_BY,ADB"`
	DeleteAT   time.Time    `col:"DELETE_AT,ADT"`
	Version    int64        `col:"VERSION,OL"`
	Name       string       `col:"NAME"`
	Password   string       `col:"PASSWORD"`
	Gender     string       `col:"GENDER"`
	Age        int          `col:"AGE"`
	Active     bool         `col:"ACTIVE"`
	SignUpTime time.Time    `col:"SIGN_UP_TIME"`
	Profile    *UserProfile `col:"PROFILE,JSON"`
	Score      float64      `col:"SCORE"`
	DOB        time.Time    `col:"DOB"`
	//Posts      []*PostRow   `col:"-,LK" ref:"ID,AUTHOR_ID" sort:"CREATE_AT DESC"`
}

func (r *UserRow) Table() (string, string, string) {
	return "FNS", "USER", "U"
}

type UserProfile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
