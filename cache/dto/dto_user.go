package dto

type UserDto struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UsersDto []UserDto
