package main

import (
	"fmt"
	"time"
	"vineapi/database"
	"vineapi/utils"
)

type JSON []byte

type Usuario struct {
	UsuarioId int            `json:"usuarioId" gorm:"primaryKey;column:id_usuario"`
	Email     string         `json:"email"`
	Senha     string         `json:"-"`
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Grupos    []*Grupoacesso `json:"grupo" gorm:"many2many:grupoacesso_usuario;foreignKey:id_usuario;joinForeignKey:id_usuario;references:id_grupoacesso;joinReferences:id_grupoacesso"`
}

type Grupoacesso struct {
	GrupoAcessoId int    `json:"grupoAcessoId" gorm:"primaryKey;column:id_grupoacesso" `
	Nome          string `json:"nome" gorm:"column:nome"`
}

type UsuarioRaw struct {
	UsuarioId int      `json:"usuarioId" gorm:"primaryKey;column:id_usuario"`
	Email     string   `json:"email"`
	Senha     string   `json:"-"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	GruposRaw JSON     `json:"-" gorm:"type:json;default:'[]'"`
	Grupos    []string `json:"grupos" gorm:"-"`
}

func queryByPreload() {
	defer utils.TimeExec(time.Now(), "preload")

	var user []Usuario

	database.GetPG().Table("tbl_usuario").Preload("Grupos").Find(&user)

}

func queryByRaw() {

	defer utils.TimeExec(time.Now(), "raw")
	users := []UsuarioRaw{}

	where, args := database.Build().Where().And("u.email  = ?", "vineboneto@gmail.com").String()

	fmt.Println(where)
	fmt.Println(args)

	sql := fmt.Sprintf(`
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
			) AS grupos_raw
		FROM tbl_usuario u %s
	`, where)

	fmt.Println(sql)

	database.GetPG().Raw(sql, args...).Scan(&users)

	fmt.Println(users)

}

func main() {

	database.Connection()

	// queryByPreload()
	queryByRaw()

}
