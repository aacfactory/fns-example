package repository

import "time"

type PostRow struct {
	Id       string            `col:"ID,pk" json:"ID"`
	User     *UserRow          `col:"USER,ref,USER_ID+ID" json:"USER"`
	CreateAT time.Time         `col:"CREATE_AT,act" json:"CREATE_AT"`
	Version  int64             `col:"VERSION,aol" json:"VERSION"`
	Title    string            `col:"TITLE" json:"TITLE"`
	Content  string            `col:"CONTENT" json:"CONTENT"`
	Comments []*PostCommentRow `col:"COMMENTS,links,ID+POST_ID,ID DESC,0:10" json:"COMMENTS"`
	Likes    int64             `col:"LIKES,vc,SELECT COUNT(1) FROM \"FNS\".\"POST_LIKE\" WHERE \"POST_ID\" = \"FNS\".\"POST\".\"ID\"" json:"LIKES"`
}

func (r PostRow) TableName() (string, string) {
	return "FNS", "POST"
}

type PostCommentRow struct {
	Id       int64     `col:"ID,incrPk" json:"ID"`
	PostId   string    `col:"POST_ID" json:"POST_ID"`
	User     *UserRow  `col:"USER,ref,USER_ID+ID" json:"USER"`
	CreateAT time.Time `col:"CREATE_AT,act" json:"CREATE_AT"`
	Content  string    `col:"CONTENT" json:"CONTENT"`
}

func (r PostCommentRow) TableName() (string, string) {
	return "FNS", "POST_COMMENT"
}

type PostLikeRow struct {
	Id     int64  `col:"ID,incrPk" json:"ID"`
	PostId string `col:"POST_ID" json:"POST_ID"`
	UserId string `col:"USER_ID" json:"USER_ID"`
}

func (r PostLikeRow) TableName() (string, string) {
	return "FNS", "POST_LIKE"
}
