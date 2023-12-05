package repositories

import (
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-contrib/databases/sql/dac"
	"github.com/aacfactory/fns/commons/times"
	"time"
)

type UserRow struct {
	dac.Audit
	Nickname string               `column:"NICKNAME" json:"NICKNAME"`
	Mobile   string               `column:"MOBILE" json:"MOBILE"`
	Gender   string               `column:"GENDER" json:"GENDER"`
	Birthday time.Time            `column:"BIRTHDAY" json:"BIRTHDAY"`
	Avatar   sql.NullJson[Avatar] `column:"AVATAR,json" json:"AVATAR_ROW" copy:"AVATAR"`
	BD       times.Date           `column:"BD" json:"BD"`
	BT       times.Time           `column:"BT" json:"BT"`
}

func (row UserRow) TableInfo() dac.TableInfo {
	return dac.Info("USER", dac.Schema("FNS"), dac.Conflicts("Id"))
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
