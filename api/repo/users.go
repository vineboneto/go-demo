package repo

import (
	"vineapi/core"
	"vineapi/database"
)

func CreateUser(user *core.CreateUserInput) int {

	database.GetPG().Table("tbl_usuario").Create(&user)

	return user.UsuarioId
}
func DeleteUser(user *core.DeleteUserInput) int {
	database.GetPG().Table("tbl_usuario").Delete(&user)

	return user.UsuarioId
}

func LoadUser(input *core.LoadUsersInput) []*core.LoadUserOutput {

	list := []*core.LoadUserOutput{}

	sql, args := database.Build().
		Raw(`
			SELECT 
				u.id_usuario, 
				u.first_name,
				u.last_name,
				u.senha,
				(
					SELECT json_agg(gr.nome) FROM tbl_grupoacesso gr
					INNER JOIN tbl_grupoacesso_usuario gu ON
						gu.id_grupoacesso = gr.id_grupoacesso AND
						gu.id_usuario = u.id_usuario
				) AS grupos
			FROM tbl_usuario u
		`).
		Where().
		And("u.id_usuario = ?", input.UsuarioId).
		AndLike("u.first_name LIKE ?", input.FirstName).
		Limit(input.Limit).
		Offset(input.Page*input.Limit - 1).
		String()

	database.GetPG().Raw(sql, args...).Scan(&list)

	return list
}

func FindByEmail(email string) *core.LoadUserOutput {
	data := &core.LoadUserOutput{}

	sql, args := database.Build().
		Raw("SELECT * FROM tbl_usuario").
		Where().
		And("email = ?", email).
		Limit(1).
		String()

	database.GetPG().Raw(sql, args...).Scan(&data)

	return data
}

func FindByID(id int) *core.LoadUserOutput {

	data := &core.LoadUserOutput{}

	sql, args := database.Build().
		Raw(`
			SELECT 
				u.id_usuario, 
				u.first_name,
				u.last_name,
				u.senha,
				(
					SELECT json_agg(gr.nome) FROM tbl_grupoacesso gr
					INNER JOIN tbl_grupoacesso_usuario gu ON
						gu.id_grupoacesso = gr.id_grupoacesso AND
						gu.id_usuario = u.id_usuario
				) AS grupos
			FROM tbl_usuario u
		`).
		Where().
		And("id_usuario = ?", id).
		Limit(1).
		String()

	database.GetPG().Raw(sql, args...).Scan(&data)

	return data
}
