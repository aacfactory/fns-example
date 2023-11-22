package users

import "time"

type User struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Age      string    `json:"age"`
	Birthday time.Time `json:"birthday"`
}
