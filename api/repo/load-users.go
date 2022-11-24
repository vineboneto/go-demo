package repo

import (
	"fmt"
	"vine-api/database"
	model "vine-api/models"
)

type QueryLoadUsers struct {
	model.User
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func CreateUser(user *model.User) int {

	fmt.Println(user)

	return 0
}

func LoadUser(query *QueryLoadUsers) []*model.User {

	list := model.ListUser{}

	s := database.Build().
		Select("*").
		From("tbl_usuario").
		Where().And("id = %d", query.Id).AndLike("first_name LIKE '%s'", query.FirstName).
		Offset(query.Page).
		Limit(query.Limit).
		String()

	database.DB.Raw(s).Scan(&list.Users)

	return list.Users
}
