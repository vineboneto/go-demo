package main

import (
	"encoding/json"
	"fmt"
	"time"
	"vineapi/database"
	"vineapi/utils"
)

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
	UsuarioId int             `json:"usuarioId" gorm:"primaryKey;column:id_usuario"`
	Email     string          `json:"email"`
	Senha     string          `json:"-"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Grupos    json.RawMessage `json:"grupos"`
}

func queryByPreload() {
	defer utils.TimeExec(time.Now(), "preload")

	var user Usuario

	database.GetPG().Table("tbl_usuario").Preload("Grupos").Where("id_usuario", 1001).First(&user)

	fmt.Println(user)

}

func queryByRaw() {

	defer utils.TimeExec(time.Now(), "raw")
	user := &UsuarioRaw{}

	database.GetPG().Raw(`
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
		WHERE u.id_usuario = 1001 LIMIT 1
	`).Scan(user)

	var grupos []string

	j, _ := user.Grupos.MarshalJSON()

	json.Unmarshal(j, &grupos)

	fmt.Println(user, grupos)
}

func main() {

	database.Connection()

	queryByPreload()
	queryByRaw()

}
