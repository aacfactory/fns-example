package posts

import (
	"github.com/aacfactory/fns-example/standalone/modules/users"
	"time"
)

// Posts
// @title post list
// @description post list
type Posts []*Post

// Post
// @title post
// @description post
type Post struct {
	// Id
	// @title id
	// @description id
	Id string `json:"id"`
	// User
	// @title author
	// @description author
	User *users.User `json:"user"`
	// CreateAT
	// @title create time
	// @description create time
	CreateAT time.Time `json:"createAt"`
	// Title
	// @title title
	// @description title
	Title string `json:"title"`
	// Content
	// @title content
	// @description content
	Content string `json:"content"`
	// Comments
	// @title comments
	// @description latest 10 comments
	Comments []*Comment `json:"comments"`
	// Likes
	// @title likes
	// @description likes num
	Likes int64 `json:"likes"`
}

// Comment
// @title comment
// @description comment
type Comment struct {
	// Id
	// @title id
	// @description id
	Id int64 `json:"id"`
	// PostId
	// @title post id
	// @description post id
	PostId string `json:"postId"`
	// User
	// @title author
	// @description author
	User *users.User `json:"user"`
	// CreateAT
	// @title create time
	// @description create time
	CreateAT time.Time `json:"createAt"`
	// Content
	// @title content
	// @description content
	Content string `json:"content"`
}
