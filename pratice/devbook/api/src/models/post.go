package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	AuthorID   uint64    `json:"author_id"`
	AuthorNick string    `json:"author_nick"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}

func (p *Post) Prepare() error {
	if erro := p.validate(); erro != nil {
		return erro
	}

	p.format()
	return nil
}

func (p *Post) validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}

	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

func (p *Post) format() {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
}
