package models

type User struct {
	Id       int    `json:"_"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
