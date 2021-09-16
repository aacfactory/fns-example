package repository

import "time"

type PostRow struct {
	Id       string            `col:"ID,PK"`
	CreateBY string            `col:"CREATE_BY,ACB"`
	CreateAT time.Time         `col:"CREATE_AT,ACT"`
	ModifyBY string            `col:"MODIFY_BY,AMB"`
	ModifyAT time.Time         `col:"MODIFY_AT,AMT"`
	Version  int64             `col:"VERSION,OL"`
	Title    string            `col:"TITLE"`
	Content  string            `col:"CONTENT"`
	Author   *UserRow          `col:"AUTHOR_ID,FK"`
	Likes    int               `col:"LIKES,VC" src:"SELECT COUNT(1) FROM \"FNS\".\"POST_LIKE\" WHERE \"POST_ID\" = \"P\".\"ID\" "`
	Comments []*PostCommentRow `col:"-,LK:SYNC" ref:"ID,POST_ID" sort:"CREATE_AT DESC"`
}

func (r *PostRow) Table() (string, string, string) {
	return "FNS", "POST", "P"
}

type PostLikeRow struct {
	Id       string    `col:"ID,PK"`
	CreateBY string    `col:"CREATE_BY,ACB"`
	CreateAT time.Time `col:"CREATE_AT,ACT"`
	Post     *PostRow  `col:"POST_ID,FK"`
}

func (r *PostLikeRow) Table() (string, string, string) {
	return "FNS", "POST_LIKE", "PL"
}

type PostCommentRow struct {
	Id       string    `col:"ID,PK"`
	CreateBY string    `col:"CREATE_BY,ACB"`
	CreateAT time.Time `col:"CREATE_AT,ACT"`
	Post     *PostRow  `col:"POST_ID,FK"`
	User     *UserRow  `col:"USER_ID,FK"`
	Content  string    `col:"CONTENT"`
}

func (r *PostCommentRow) Table() (string, string, string) {
	return "FNS", "POST_COMMENT", "PC"
}
