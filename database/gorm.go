package main

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StageFluxoPag struct {
	Id     int
	Codemp int
	Codfil int
}

type Cliente struct {
	Codcli int
	Nomcli string
}

func timeTrack(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}

func main() {

	defer timeTrack(time.Now())

	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1234 dbname=erpweb port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	var stage StageFluxoPag

	// var stages []StageFluxoPag

	db.Select("id, codemp, codfil").Where("id = ?", 280350).Table("stage_fluxo_pag").First(&stage)

	// db.Select("*").Table("stage_fluxo_pag").Limit(500).Find(&stages)

	fmt.Println(stage)

	// db2, _ := gorm.Open(sqlserver.Open("url"), &gorm.Config{})

	// var clientes []Cliente

	// db2.Select("codcli, nomcli").Table("e085cli").Find(&clientes)

	// fmt.Println(len(clientes))
}
