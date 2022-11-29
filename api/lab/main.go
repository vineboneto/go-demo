package main

import (
	"fmt"
	"time"
	"vineapi/database"
	"vineapi/utils"

	"github.com/joho/godotenv"
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
	UsuarioId int    `json:"usuarioId" gorm:"primaryKey;column:id_usuario"`
	Email     string `json:"email"`
	Senha     string `json:"-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Grupos    []byte `json:"grupos" gorm:"-"`
}

func queryByPreload() {
	defer utils.TimeExec(time.Now(), "preload")

	var user []Usuario

	database.GetPG().Table("tbl_usuario").Preload("Grupos").Find(&user)

	fmt.Println(len(user))

}

func queryByRaw() {

	defer utils.TimeExec(time.Now(), "raw")
	users := []UsuarioRaw{}

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
		String()

	database.GetPG().Raw(sql, args...).Scan(&users)

	fmt.Println(len(users))

}

func main() {

	godotenv.Load()

	database.Connection()

	for i := 0; i < 100; i++ {
		queryByPreload()
		queryByRaw()
	}

}
