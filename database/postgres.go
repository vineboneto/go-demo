package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Client struct {
	ID   int
	Name string
}

type List struct {
	Clients []*Client
}

func timeTrack(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}

func main() {

	defer timeTrack(time.Now())

	err := godotenv.Load()

	conn, err := sql.Open("postgres", os.Getenv("DATABASE_URL_POSTGRES"))

	defer conn.Close()

	if err != nil {
		fmt.Println(err)
	}

	rows, err := conn.Query(`SELECT nomcli, codcli FROM tbl_cliente`)

	if err != nil {
		fmt.Println("Entrei aqui")
		fmt.Println(err)
	}

	defer rows.Close()

	list := &List{}

	for rows.Next() {

		var codcli int
		var nomcli string

		err := rows.Scan(&nomcli, &codcli)

		if err != nil {
			fmt.Println(err)
		}

		client := &Client{ID: codcli, Name: nomcli}

		list.Clients = append(list.Clients, client)
	}

	fmt.Println(len(list.Clients))

}
