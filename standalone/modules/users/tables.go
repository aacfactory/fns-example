package users

import (
	"github.com/aacfactory/json"
	"time"
)

type UserRow struct {
	Id         string          `col:"ID"`
	Name       string          `col:"NAME"`
	Password   string          `col:"PASSWORD"`
	Gender     string          `col:"GENDER"`
	Age        int             `col:"AGE"`
	Active     bool            `col:"ACTIVE"`
	SignUpTime time.Time       `col:"SIGN_UP_TIME"`
	Profile    json.RawMessage `col:"PROFILE"`
	Score      float64         `col:"SCORE"`
	DOB        time.Time       `col:"DOB"`
}
