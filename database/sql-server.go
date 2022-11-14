package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/denisenkom/go-mssqldb/azuread"
	"github.com/joho/godotenv"
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

func handlerError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	defer timeTrack(time.Now())

	err := godotenv.Load()

	handlerError(err)

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(os.Getenv("MSSQL_USER"), os.Getenv("MSSQL_PASSWORD")),
		Host:   fmt.Sprintf("%s:%s", os.Getenv("MSSQL_HOST"), os.Getenv("MSSQL_PORT")),
	}

	conn, err := sql.Open(azuread.DriverName, u.String())

	handlerError(err)

	defer conn.Close()

	rows, err := conn.Query(`SELECT TOP 1 nomcli, codcli from sapiens.dbo.e085cli`)

	handlerError(err)

	defer rows.Close()

	list := &List{}

	for rows.Next() {
		var codcli int
		var nomcli string

		err := rows.Scan(&nomcli, &codcli)

		handlerError(err)

		cliente := &Client{Name: nomcli, ID: codcli}

		list.Clients = append(list.Clients, cliente)
	}

	fmt.Println(len(list.Clients))
}
