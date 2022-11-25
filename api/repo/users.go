package repo

import (
	"encoding/json"
	"fmt"
	"vineapi/database"
)

type LoadUserOutput struct {
	Id        int             `json:"id"`
	Email     string          `json:"email"`
	Senha     string          `json:"-"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Grupos    json.RawMessage `json:"grupos"`
}

type LoadUsersInput struct {
	Id        int    `json:"id" form:"id"`
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
	Id        int    `json:"id"`
	Email     string `json:"email" binding:"required,email"`
	Senha     string `json:"senha" binding:"required"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
}

type DeleteUserInput struct {
	Id int `json:"id"`
}

func CreateUser(user *CreateUserInput) int {

	database.DB.Table("tbl_usuario").Create(&user)

	return user.Id
}

func DeleteUser(user *DeleteUserInput) int {
	database.DB.Table("tbl_usuario").Delete(&user)

	return user.Id
}

func LoadUser(input *LoadUsersInput) []*LoadUserOutput {

	list := []*LoadUserOutput{}

	where := database.Build().Where().And("u.id = %d", input.Id).String()

	limitOffset := database.Build().Limit(input.Limit).Offset(input.Page).String()

	s := fmt.Sprintf(`
		SELECT 
			u.id, 
			u.first_name,
			u.last_name,
			(
				SELECT json_agg(gr.nome) FROM tbl_grupoacesso gr
				INNER JOIN tbl_grupoacesso_usuario gu ON gu.id_grupoacesso = gr.id AND gu.id_usuario = u.id
			) AS grupos
		FROM tbl_usuario u
		%s
		%s
	`, where, limitOffset)

	database.DB.Raw(s).Scan(&list)

	return list
}

func FindByEmail(email string) *LoadUserOutput {
	data := &LoadUserOutput{}

	database.DB.
		Raw(fmt.Sprintf("SELECT * FROM tbl_usuario WHERE email = '%s' LIMIT 1", email)).
		Scan(&data)

	return data
}

func FindByID(id int) *LoadUserOutput {
	data := &LoadUserOutput{}

	database.DB.
		Raw(fmt.Sprintf(`
			SELECT 
				u.id, 
				u.first_name,
				u.last_name,
				(
					SELECT json_agg(gr.nome) FROM tbl_grupoacesso gr
					INNER JOIN tbl_grupoacesso_usuario gu ON gu.id_grupoacesso = gr.id AND gu.id_usuario = u.id
				) AS grupos
			FROM tbl_usuario u
			WHERE u.id = %d LIMIT 1
		`, id)).
		Scan(data)

	return data
}
