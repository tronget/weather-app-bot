package models

type User struct {
	Username string
	LangCode string
}

func NewUser(username, langCode string) *User {
	return &User{username, langCode}
}
