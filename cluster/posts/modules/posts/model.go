package posts

import "time"

type Posts []Post

type Post struct {
	Id       string    `json:"id"`
	UserId   string    `json:"userId"`
	CreateAT time.Time `json:"createAT"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Likes    int64     `json:"likes"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id       int64     `json:"id"`
	PostId   string    `json:"postId"`
	UserId   string    `json:"userId"`
	CreateAT time.Time `json:"createAT"`
	Content  string    `json:"content"`
}
