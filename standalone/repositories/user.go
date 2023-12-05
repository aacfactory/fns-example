package repositories

import (
	"github.com/aacfactory/fns-contrib/databases/sql"
	"github.com/aacfactory/fns-contrib/databases/sql/dac"
	"github.com/aacfactory/fns/commons/times"
	"time"
)

type UserRow struct {
	dac.Audit
	Nickname string               `column:"NICKNAME"`
	Mobile   string               `column:"MOBILE"`
	Gender   string               `column:"GENDER"`
	Birthday time.Time            `column:"BIRTHDAY"`
	Avatar   sql.NullJson[Avatar] `column:"AVATAR,json"`
	BD       times.Date           `column:"BD"`
	BT       times.Time           `column:"BT"`
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

type UserGenderCount struct {
	Gender string `column:"GENDER" json:"gender"`
	Count  int64  `column:"ID,vc,agg,COUNT" json:"count"`
}

func (u UserGenderCount) ViewInfo() dac.ViewInfo {
	return dac.TableView(UserRow{})
}
