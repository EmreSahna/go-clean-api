package model

type User struct {
	Base
	Username string `json:"username"`
	Password string `json:"password"`
}
