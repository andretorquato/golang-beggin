package models

type Password struct {
	NewPass string `json:"new_password"`
	OldPass string `json:"old_password"`
}
