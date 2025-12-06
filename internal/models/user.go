package models

type User struct {
	Id       int    `json:"_"`
	Username string `json:"username"`
	Password string `json:"password"`
}
