package repository

import "time"

type PostRow struct {
	Id       string            `col:"ID,PK"`
	CreateBY string            `col:"CREATE_BY,CREATE_BY"`
	CreateAT time.Time         `col:"CREATE_AT,CREATE_AT"`
	ModifyBY string            `col:"MODIFY_BY,MODIFY_BY"`
	ModifyAT time.Time         `col:"MODIFY_AT,MODIFY_AT"`
	Version  int64             `col:"VERSION,VERSION"`
	Title    string            `col:"TITLE"`
	Content  string            `col:"CONTENT"`
	Author   *UserRow          `col:"AUTHOR_ID,FK:SYNC"`
	Likes    int               `col:"LIKES,VC" src:"SELECT COUNT(1) FROM \"FNS\".\"POST_LIKE\" WHERE \"POST_ID\" = POST_ID = \"P\".\"ID\" "`
	Comments []*PostCommentRow `col:"-,LK:SYNC" sort:"CREATE_AT DESC"`
}

func (r *PostRow) Table() (string, string, string) {
	return "FNS", "POST", "P"
}

type PostLikeRow struct {
	Id       string    `col:"ID,PK"`
	CreateBY string    `col:"CREATE_BY,CREATE_BY"`
	CreateAT time.Time `col:"CREATE_AT,CREATE_AT"`
	Post     *PostRow  `col:"POST_ID,FK"`
}

func (r *PostLikeRow) Table() (string, string, string) {
	return "FNS", "POST_LIKE", "PL"
}

type PostCommentRow struct {
	Id       string    `col:"ID,PK"`
	CreateBY string    `col:"CREATE_BY,CREATE_BY"`
	CreateAT time.Time `col:"CREATE_AT,CREATE_AT"`
	Post     *PostRow  `col:"POST_ID,FK"`
	User     *UserRow  `col:"USER_ID,FK"`
	Content  string    `col:"CONTENT"`
}

func (r *PostCommentRow) Table() (string, string, string) {
	return "FNS", "POST_COMMENT", "PC"
}
