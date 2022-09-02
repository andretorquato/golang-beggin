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

func (repo Posts) GetAll(userID uint64) ([]models.Post, error) {
	rows, erro := repo.db.Query(`
	SELECT DISTINCT p.*, u.nick FROM posts p
	INNER JOIN users u ON u.id = p.author_id 
	INNER JOIN followers f ON f.user_id = p.author_id  
	WHERE u.id = ? OR f.follower_id = ?
	ORDER BY 1 DESC`,
		userID, userID)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if erro := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorNick); erro != nil {
			return nil, erro
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (repo Posts) Update(postID uint64, post models.Post) error {
	statement, erro := repo.db.Prepare("UPDATE posts SET title = ?, content = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(post.Title, post.Content, postID); erro != nil {
		return erro
	}

	return nil
}

func (repo Posts) Delete(postID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM posts WHERE id = ?")
	if erro != nil {
		return nil
	}
	defer statement.Close()

	if _, erro = statement.Exec(postID); erro != nil {
		return erro
	}

	return nil
}
