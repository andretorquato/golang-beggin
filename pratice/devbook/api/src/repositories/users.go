package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (uint64, error) {
	statement, erro := repo.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastIDInserted, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastIDInserted), nil
}

func (repo Users) Search(query string) ([]models.User, error) {
	newQuery := fmt.Sprintf("%%%s%%", query)
	statement, erro := repo.db.Prepare("SELECT id, name, nick, email, created_at FROM users WHERE name LIKE ? OR nick LIKE ?")
	if erro != nil {
		return nil, erro
	}
	defer statement.Close()
	rows, erro := statement.Query(newQuery, newQuery)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		var user models.User
		erro := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	if erro = rows.Err(); erro != nil {
		return nil, erro
	}
	return users, nil
}

func (repo Users) FindByID(id uint64) (models.User, error) {
	statement, erro := repo.db.Prepare("SELECT id, name, nick, email, created_at FROM users WHERE id = ?")
	if erro != nil {
		return models.User{}, erro
	}
	defer statement.Close()
	row := statement.QueryRow(id)
	var user models.User
	erro = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
	if erro != nil {
		return models.User{}, erro
	}
	return user, nil
}

func (repo Users) Update(user models.User, id uint64) error {
	statement, erro := repo.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(user.Name, user.Nick, user.Email, id)
	if erro != nil {
		return erro
	}
	return nil
}

func (repo Users) Delete(id uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM users WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (repo Users) FindByEmail(email string) (models.User, error) {
	statement, erro := repo.db.Prepare("SELECT id, password FROM users WHERE email = ?")
	if erro != nil {
		return models.User{}, nil
	}
	defer statement.Close()
	row := statement.QueryRow(email)
	var user models.User
	erro = row.Scan(&user.ID, &user.Password)
	if erro != nil {
		return models.User{}, nil
	}
	return user, nil
}

func (repo Users) Follow(userID, followID uint64) error {
	statement, erro := repo.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES(?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(userID, followID)
	if erro != nil {
		return erro
	}
	return nil
}

func (repo Users) UnFollow(userID, unFollowerID uint64) error {
	statement, erro := repo.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	_, erro = statement.Exec(userID, unFollowerID)
	if erro != nil {
		return erro
	}
	return nil
}

func (repo Users) GetFollowers(userID uint64) ([]models.User, error) {
	rows, erro := repo.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.created_at 
	FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ?`,
		userID)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		erro := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo Users) GetFollowing(userID uint64) ([]models.User, error) {
	rows, erro := repo.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.created_at
	FROM users u INNER JOIN followers f ON u.id = f.user_id WHERE f.follower_id = ?
	`, userID)
	if erro != nil {
		return nil, erro
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		erro := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}

	return users, nil
}
