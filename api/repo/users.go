package repo

import (
	"fmt"
	"vineapi/core"
	"vineapi/database"
)

func CreateUser(user *core.CreateUserInput) int {

	database.GetPG().Table("tbl_usuario").Create(&user)

	return user.Id
}
func DeleteUser(user *core.DeleteUserInput) int {
	database.GetPG().Table("tbl_usuario").Delete(&user)

	return user.Id
}

func LoadUser(input *core.LoadUsersInput) []*core.LoadUserOutput {

	list := []*core.LoadUserOutput{}

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

	database.GetPG().Raw(s).Scan(&list)

	return list
}

func FindByEmail(email string) *core.LoadUserOutput {
	data := &core.LoadUserOutput{}

	database.GetPG().
		Raw(fmt.Sprintf("SELECT * FROM tbl_usuario WHERE email = '%s' LIMIT 1", email)).
		Scan(&data)

	return data
}

func FindByID(id int) *core.LoadUserOutput {
	data := &core.LoadUserOutput{}

	database.GetPG().
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
