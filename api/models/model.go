package model

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}

type ListUser struct {
	Users []*User
}
