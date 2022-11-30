package models

type User struct {
	ID       int8   `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Name     string `db:"name" json:"name"`
}

type UserLoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
