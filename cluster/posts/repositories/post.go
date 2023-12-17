package repositories

import (
	"github.com/aacfactory/fns-contrib/databases/sql/dac"
	"time"
)

type PostRow struct {
	Id       string          `column:"ID,pk"`
	UserId   string          `column:"USER_ID"`
	CreateAT time.Time       `column:"CREATE_AT,act"`
	Version  int64           `column:"VERSION,aol"`
	Title    string          `column:"TITLE"`
	Content  string          `column:"CONTENT"`
	Comments PostCommentRows `column:"COMMENTS,links,Id+PostId,orders:Id@desc,length:10"`
	Likes    int64           `column:"LIKES,vc,basic,SELECT COUNT(1) FROM \"FNS\".\"POST_LIKE\" WHERE \"POST_ID\" = \"FNS\".\"POST\".\"ID\""`
}

func (row PostRow) TableInfo() dac.TableInfo {
	return dac.Info("POST", dac.Schema("FNS"))
}

type PostCommentRows []PostCommentRow

type PostCommentRow struct {
	Id       int64     `column:"ID,pk,incr"`
	PostId   string    `column:"POST_ID"`
	UserId   string    `column:"USER_ID"`
	CreateAT time.Time `column:"CREATE_AT,act"`
	Content  string    `column:"CONTENT"`
}

func (row PostCommentRow) TableInfo() dac.TableInfo {
	return dac.Info("POST_COMMENT", dac.Schema("FNS"))
}

type PostLikeRow struct {
	Id     int64  `column:"ID,pk,incr"`
	PostId string `column:"POST_ID"`
	UserId string `column:"USER_ID"`
}

func (row PostLikeRow) TableInfo() dac.TableInfo {
	return dac.Info("POST_LIKE", dac.Schema("FNS"))
}
