package users

import (
	"github.com/aacfactory/fns-example/cluster/users/repository"
	"github.com/aacfactory/json"
	"time"
)

// User
// @title User
// @description User model
type User struct {
	// Id
	// @title Id
	// @description Id
	Id string `json:"id"`
	// CreateAT
	// @title create time
	// @description create time
	CreateAT time.Time `json:"createAt"`
	// Nickname
	// @title nickname
	// @description nickname
	Nickname string `json:"nickname"`
	// Mobile
	// @title mobile
	// @description mobile
	Mobile string `json:"mobile"`
	// Gender
	// @title gender
	// @enum F(female),M(male),N(unknown)
	// @description gender
	Gender string `json:"gender"`
	// Birthday
	// @title birthday
	// @description birthday
	Birthday json.Date `json:"birthday"`
	// Avatar
	// @title user avatar
	// @description user avatar
	Avatar *repository.Avatar `json:"avatar"`
}
