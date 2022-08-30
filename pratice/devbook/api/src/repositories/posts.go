package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repo Posts) Create(post models.Post) (uint64, error) {
	statement, erro := repo.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES(?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(post.Title, post.Content, post.AuthorID)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}

func (repo Posts) GetByID(postID uint64) (models.Post, error) {
	row, erro := repo.db.Query(`SELECT p.*, u.nick FROM posts p INNER JOIN users u ON u.id = p.author_id WHERE p.id = ?`, postID)
	if erro != nil {
		return models.Post{}, erro
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		fmt.Println(row.Scan())
		if erro := row.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); erro != nil {
			return models.Post{}, erro
		}
	}

	return post, nil
}
