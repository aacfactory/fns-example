package repositories

import (
	"github.com/aacfactory/fns-contrib/databases/sql/dac"
	"time"
)

type PostRow struct {
	Id       string            `column:"ID,pk" json:"ID"`
	User     *UserRow          `column:"USER_ID,ref,User+Id" json:"USER"`
	CreateAT time.Time         `column:"CREATE_AT,act" json:"CREATE_AT"`
	Version  int64             `column:"VERSION,aol" json:"VERSION"`
	Title    string            `column:"TITLE" json:"TITLE"`
	Content  string            `column:"CONTENT" json:"CONTENT"`
	Comments []*PostCommentRow `column:"COMMENTS,links,Id+PostId,orders:Id@desc,length:10" json:"COMMENTS"`
	Likes    int64             `column:"LIKES,vc,basic,SELECT COUNT(1) FROM \"FNS\".\"POST_LIKE\" WHERE \"POST_ID\" = \"FNS\".\"POST\".\"ID\"" json:"LIKES"`
}

func (row PostRow) TableInfo() dac.TableInfo {
	return dac.Info("POST", dac.Schema("FNS"))
}

type PostCommentRow struct {
	Id       int64     `column:"ID,pk,incr" json:"ID"`
	PostId   string    `column:"POST_ID" json:"POST_ID"`
	User     *UserRow  `column:"USER_ID,ref,User+Id" json:"USER"`
	CreateAT time.Time `column:"CREATE_AT,act" json:"CREATE_AT"`
	Content  string    `column:"CONTENT" json:"CONTENT"`
}

func (row PostCommentRow) TableInfo() dac.TableInfo {
	return dac.Info("POST_COMMENT", dac.Schema("FNS"))
}

type PostLikeRow struct {
	Id     int64  `column:"ID,pk,incr" json:"ID"`
	PostId string `column:"POST_ID" json:"POST_ID"`
	UserId string `column:"USER_ID" json:"USER_ID"`
}

func (row PostLikeRow) TableInfo() dac.TableInfo {
	return dac.Info("POST_LIKE", dac.Schema("FNS"))
}
