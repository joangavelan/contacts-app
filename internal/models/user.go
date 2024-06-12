package models

type User struct {
	Id       int64
	Username string
	Email    string
	Password string
}

type UserContext struct {
	Id       int64
	Username string
}
