package repository

import "time"

type UserRow struct {
	Id       string    `col:"ID" json:"ID"`
	CreateBY string    `col:"CREATE_BY,acb" json:"CREATE_BY"`
	CreateAT time.Time `col:"CREATE_AT,act" json:"CREATE_AT"`
	ModifyBY string    `col:"MODIFY_BY,amb" json:"MODIFY_BY"`
	ModifyAT time.Time `col:"MODIFY_AT,amt" json:"MODIFY_AT"`
	DeleteBY string    `col:"DELETE_BY,adb" json:"DELETE_BY"`
	DeleteAT time.Time `col:"DELETE_AT,adt" json:"DELETE_AT"`
	Version  int64     `col:"VERSION,aol" json:"VERSION"`
	Nickname string    `col:"NICKNAME" json:"NICKNAME"`
	Mobile   string    `col:"MOBILE" json:"MOBILE"`
	Gender   string    `col:"GENDER" json:"GENDER"`
	Birthday time.Time `col:"BIRTHDAY" json:"BIRTHDAY"`
	Avatar   *Avatar   `col:"AVATAR,json" json:"AVATAR_ROW" copy:"AVATAR"`
}

func (r UserRow) TableName() (string, string) {
	return "FNS", "USER"
}

// Avatar
// @title Avatar
// @description Avatar info
type Avatar struct {
	// Schema
	// @title http schema
	// @description http schema
	Schema string `json:"schema"`
	// Domain
	// @title domain
	// @description domain
	Domain string `json:"domain"`
	// Path
	// @title uri path
	// @description uri path
	Path string `json:"path"`
	// MimeType
	// @title mime type
	// @description mime type
	MimeType string `json:"mimeType"`
	// URL
	// @title url
	// @description full url
	URL string `json:"url"`
}
