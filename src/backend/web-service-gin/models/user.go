package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // In der Realität sollte das Passwort gehasht sein
}
