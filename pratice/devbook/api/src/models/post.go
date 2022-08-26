package models

import "time"

type Post struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint64    `json:"author_id"`
	AuthorNick string    `json:"author_nick"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
