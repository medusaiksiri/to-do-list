package models

type User struct {
	ID       uint
	Username string
	Password string
	UserType string // "normal" or "admin"
}
