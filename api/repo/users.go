package repo

import (
	"fmt"
	"vineapi/database"
)

type LoadUserOutput struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Senha     string `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type LoadUsersInput struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
}

type FindEmailInput struct {
	Email string
}

type CreateUserInput struct {
	Id        int    `json:"id"`
	Email     string `json:"email" validate:"required,email"`
	Senha     string `json:"senha" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
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

	s := database.Build().
		Select("*").
		From("tbl_usuario").
		Where().And("id = %d", input.Id).AndLike("first_name LIKE '%s'", input.FirstName).
		Offset(input.Page).
		Limit(input.Limit).
		String()

	database.DB.Raw(s).Scan(&list)

	return list
}

func FindEmail(email string) *LoadUserOutput {
	data := &LoadUserOutput{}

	database.DB.
		Raw(fmt.Sprintf("SELECT * FROM tbl_usuario WHERE email = '%s' LIMIT 1", email)).
		Scan(&data)

	return data
}
