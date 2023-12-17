package repositories

import (
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-contrib/databases/sql/dac"
	"time"
)

type UserRow struct {
	dac.Audit
	Nickname string               `column:"NICKNAME" json:"nickname"`
	Mobile   string               `column:"MOBILE" json:"mobile"`
	Gender   string               `column:"GENDER" json:"gender"`
	Birthday time.Time            `column:"BIRTHDAY" json:"birthday"`
	Avatar   sql.NullJson[Avatar] `column:"AVATAR,json" json:"avatar"`
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
