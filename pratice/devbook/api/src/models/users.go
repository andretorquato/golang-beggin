package models

import (
	"api/src/security"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Nick      string `json:"nick,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	if erro := user.formatter(step); erro != nil {
		return erro
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Nick == "" {
		return errors.New("nick is required")
	}

	if user.Email == "" {
		return errors.New("e-mail is required")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("e-mail is invalid")
	}

	if step == "create" && user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func (user *User) formatter(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if step == "create" {
		hash, erro := security.Hash(user.Password)
		if erro != nil {
			return erro
		}
		user.Password = string(hash)
	}

	return nil
}
