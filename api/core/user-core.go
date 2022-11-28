package core

import "encoding/json"

type LoadUserOutput struct {
	IdUsuario int             `json:"usuarioId"`
	Email     string          `json:"email"`
	Senha     string          `json:"-"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Grupos    json.RawMessage `json:"grupos"`
}

type LoadUsersInput struct {
	IdUsuario int    `json:"id" form:"idUsuario"`
	Email     string `json:"email" form:"email"`
	FirstName string `json:"firstName" form:"firstName"`
	LastName  string `json:"lastName" form:"lastName"`
	Page      int    `json:"page" form:"page"`
	Limit     int    `json:"limit" form:"limit"`
}

type FindEmailInput struct {
	Email string
}

type CreateUserInput struct {
	IdUsuario int    `json:"usuarioId" gorm:"primaryKey;column:id_usuario"`
	Email     string `json:"email" binding:"required,email"`
	Senha     string `json:"senha" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type DeleteUserInput struct {
	IdUsuario int `json:"usuarioId"`
}
